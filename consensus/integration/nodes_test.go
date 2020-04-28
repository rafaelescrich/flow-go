package integration_test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/dapperlabs/flow-go/cmd/bootstrap/run"
	"github.com/dapperlabs/flow-go/consensus"
	"github.com/dapperlabs/flow-go/consensus/hotstuff"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/helper"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications/pubsub"
	"github.com/dapperlabs/flow-go/engine/common/synchronization"
	"github.com/dapperlabs/flow-go/engine/consensus/compliance"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/flow/filter"
	"github.com/dapperlabs/flow-go/module/buffer"
	builder "github.com/dapperlabs/flow-go/module/builder/consensus"
	finalizer "github.com/dapperlabs/flow-go/module/finalizer/consensus"
	"github.com/dapperlabs/flow-go/module/local"
	"github.com/dapperlabs/flow-go/module/mempool/stdmap"
	module "github.com/dapperlabs/flow-go/module/mock"
	networkmock "github.com/dapperlabs/flow-go/network/mock"
	protocol "github.com/dapperlabs/flow-go/state/protocol/badger"
	storage "github.com/dapperlabs/flow-go/storage/badger"
	"github.com/dapperlabs/flow-go/utils/unittest"
	"github.com/dgraph-io/badger/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const hotstuffTimeout = 200 * time.Millisecond

type Node struct {
	db            *badger.DB
	dbDir         string
	index         int
	log           zerolog.Logger
	id            *flow.Identity
	compliance    *compliance.Engine
	sync          *synchronization.Engine
	hot           *hotstuff.EventLoop
	state         *protocol.State
	headers       *storage.Headers
	views         *storage.Views
	net           *Network
	blockproposal int
	blockvote     int
	syncreq       int
	syncresp      int
	rangereq      int
	batchreq      int
	batchresp     int
	sync.Mutex
}

func createNodes(t *testing.T, n int, stopAtView uint64) ([]*Node, *Stopper) {
	// create n consensus nodes
	participants := make([]*flow.Identity, 0)
	for i := 0; i < n; i++ {
		identity := unittest.IdentityFixture()
		participants = append(participants, identity)
	}

	// create non-consensus nodes
	collection := unittest.IdentityFixture()
	collection.Role = flow.RoleCollection

	verification := unittest.IdentityFixture()
	verification.Role = flow.RoleVerification

	execution := unittest.IdentityFixture()
	execution.Role = flow.RoleExecution

	allParitipants := append(participants, collection, verification, execution)

	// add all identities to genesis block and
	// create and bootstrap consensus node with the genesis
	genesis := run.GenerateRootBlock(allParitipants, run.GenerateRootSeal([]byte{}))

	stopper := NewStopper(stopAtView)
	nodes := make([]*Node, 0, len(participants))
	for i, identity := range participants {
		node := createNode(t, i, identity, participants, &genesis, stopper)
		nodes = append(nodes, node)
	}

	return nodes, stopper
}

func createNode(t *testing.T, index int, identity *flow.Identity, participants flow.IdentityList, genesis *flow.Block, stopper *Stopper) *Node {
	db, dbDir := unittest.TempBadgerDB(t)
	state, err := protocol.NewState(db)
	require.NoError(t, err)

	err = state.Mutate().Bootstrap(genesis)
	require.NoError(t, err)

	localID := identity.ID()

	node := &Node{
		db:    db,
		dbDir: dbDir,
		index: index,
		id:    identity,
	}

	stopConsumer := stopper.AddNode(node)

	// log with node index
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	// log := zerolog.New(ioutil.Discard).Level(zerolog.DebugLevel).With().Timestamp().Int("index", index).Hex("local_id", localID[:]).Logger()
	log := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Int("index", index).Hex("local_id", localID[:]).Logger()
	notifier := notifications.NewLogConsumer(log)
	dis := pubsub.NewDistributor()
	dis.AddConsumer(stopConsumer)
	dis.AddConsumer(notifier)

	// initialize no-op metrics mock
	metrics := &module.Metrics{}
	metrics.On("HotStuffBusyDuration", mock.Anything, mock.Anything)
	metrics.On("HotStuffIdleDuration", mock.Anything, mock.Anything)
	metrics.On("HotStuffWaitDuration", mock.Anything, mock.Anything)

	// make local
	priv := helper.MakeBLSKey(t)
	local, err := local.New(identity, priv)
	require.NoError(t, err)

	// make network
	net := NewNetwork()

	headersDB := storage.NewHeaders(db)
	payloadsDB := storage.NewPayloads(db)
	blocksDB := storage.NewBlocks(db)
	viewsDB := storage.NewViews(db)

	guarantees, err := stdmap.NewGuarantees(100000)
	require.NoError(t, err)
	seals, err := stdmap.NewSeals(100000)
	require.NoError(t, err)

	// initialize the block builder
	build := builder.NewBuilder(db, guarantees, seals)

	signer := &Signer{identity.ID()}

	// initialize the pending blocks cache
	cache := buffer.NewPendingBlocks()

	rootHeader := &genesis.Header
	rootQC := &model.QuorumCertificate{
		View:      genesis.View,
		BlockID:   genesis.ID(),
		SignerIDs: nil, // TODO
		SigData:   nil,
	}
	selector := filter.HasRole(flow.RoleConsensus)

	// initialize the block finalizer
	final := finalizer.NewFinalizer(db, guarantees, seals)

	prov := &networkmock.Engine{}

	// initialize the compliance engine
	comp, err := compliance.New(log, net, local, state, headersDB, payloadsDB, prov, cache)
	require.NoError(t, err)

	// initialize the synchronization engine
	sync, err := synchronization.New(log, net, local, state, blocksDB, comp)
	require.NoError(t, err)

	// initialize the block finalizer
	hot, err := consensus.NewParticipant(log, dis, metrics, headersDB,
		viewsDB, state, local, build, final, signer, comp, selector, rootHeader,
		rootQC, consensus.WithTimeout(hotstuffTimeout))

	require.NoError(t, err)

	comp = comp.WithSynchronization(sync).WithConsensus(hot)

	node.compliance = comp
	node.sync = sync
	node.state = state
	node.hot = hot
	node.headers = headersDB
	node.views = viewsDB
	node.net = net
	node.log = log

	return node
}
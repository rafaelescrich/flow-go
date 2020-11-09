package backend

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/utils/unittest"
	accessproto "github.com/onflow/flow/protobuf/go/flow/access"
	"github.com/onflow/flow/protobuf/go/flow/entities"
)

// TestHistoricalTransaction tests to see if the historical transaction status can be retrieved
func (suite *Suite) TestHistoricalTransaction() {

	ctx := context.Background()
	collection := unittest.CollectionFixture(1)
	transactionBody := collection.Transactions[0]

	txID := transactionBody.ID()
	// transaction storage returns the corresponding transaction
	suite.transactions.
		On("ByID", txID).
		Return(nil, status.Errorf(codes.NotFound, "not found on main node"))

	accessEventReq := accessproto.GetTransactionRequest{
		Id: txID[:],
	}

	accessEventResp := accessproto.TransactionResultResponse{
		Status: entities.TransactionStatus(flow.TransactionStatusSealed),
		Events: nil,
	}

	backend := New(
		suite.state,
		suite.execClient,
		nil,
		[]accessproto.AccessAPIClient{suite.historicalAccessClient},
		suite.blocks,
		suite.headers,
		suite.collections,
		suite.transactions,
		suite.chainID,
		metrics.NewNoopCollector(),
		0,
		nil,
		false,
	)

	// Successfully return empty event list
	suite.historicalAccessClient.
		On("GetTransactionResult", ctx, &accessEventReq).
		Return(&accessEventResp, nil).
		Once()

	// first call - when block under test is greater height than the sealed head, but execution node does not know about Tx
	result, err := backend.GetTransactionResult(ctx, txID)
	suite.checkResponse(result, err)

	// status should be finalized since the sealed blocks is smaller in height
	suite.Assert().Equal(flow.TransactionStatusSealed, result.Status)

	suite.assertAllExpectations()
}
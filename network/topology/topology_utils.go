package topology

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/onflow/flow-go-sdk/crypto"

	"github.com/onflow/flow-go/model/encoding"
	"github.com/onflow/flow-go/model/flow"
)

// FanoutFunc represents a function type that receiving total number of nodes
// in flow system, returns fanout of individual nodes.
type FanoutFunc func(size int) int

// LinearFanoutFunc guarantees full network connectivity in a deterministic way.
// Given system of `size` nodes, it returns `size+1/2`.
func LinearFanout(size int) int {
	fanout := math.Ceil(float64(size+1) / 2)
	return int(fanout)
}

// intSeedFromID generates a int64 seed from a flow.Identifier.
func intSeedFromID(id flow.Identifier) (int64, error) {
	var seed int64
	buf := bytes.NewBuffer(id[:])
	if err := binary.Read(buf, binary.LittleEndian, &seed); err != nil {
		return -1, fmt.Errorf("could not read random bytes: %w", err)
	}
	return seed, nil
}

// byteSeedFromID generates a byte seed from a flow.Identifier.
func byteSeedFromID(id flow.Identifier) ([]byte, error) {
	h, err := crypto.NewHasher(crypto.SHA3_256)
	if err != nil {
		return nil, fmt.Errorf("could not generate hasher: %w", err)
	}

	encodedId, err := encoding.DefaultEncoder.Encode(id)
	if err != nil {
		return nil, fmt.Errorf("could not encode id: %w", err)
	}

	return h.ComputeHash(encodedId), nil
}

package convert

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"

	"github.com/dapperlabs/flow/protobuf/go/flow/entities"

	"github.com/dapperlabs/flow-go/model/flow"
)

func MessageToAccountSignature(m *entities.AccountSignature) flow.AccountSignature {
	return flow.AccountSignature{
		Account:   flow.BytesToAddress(m.GetAccount()),
		Signature: m.GetSignature(),
	}
}

func AccountSignatureToMessage(a flow.AccountSignature) *entities.AccountSignature {
	return &entities.AccountSignature{
		Account:   a.Account.Bytes(),
		Signature: a.Signature,
	}
}

func MessageToTransaction(m *entities.Transaction) (flow.TransactionBody, error) {
	if m == nil {
		return flow.TransactionBody{}, fmt.Errorf("message is empty")
	}

	scriptAccounts := make([]flow.Address, len(m.ScriptAccounts))
	for i, account := range m.ScriptAccounts {
		scriptAccounts[i] = flow.BytesToAddress(account)
	}

	signatures := make([]flow.AccountSignature, len(m.Signatures))
	for i, accountSig := range m.Signatures {
		signatures[i] = MessageToAccountSignature(accountSig)
	}

	return flow.TransactionBody{
		Script:           m.GetScript(),
		ReferenceBlockID: flow.HashToID(m.ReferenceBlockId),
		PayerAccount:     flow.BytesToAddress(m.PayerAccount),
		ScriptAccounts:   scriptAccounts,
		Signatures:       signatures,
	}, nil
}

func TransactionToMessage(t flow.TransactionBody) *entities.Transaction {
	scriptAccounts := make([][]byte, len(t.ScriptAccounts))
	for i, account := range t.ScriptAccounts {
		scriptAccounts[i] = account.Bytes()
	}

	signatures := make([]*entities.AccountSignature, len(t.Signatures))
	for i, accountSig := range t.Signatures {
		signatures[i] = AccountSignatureToMessage(accountSig)
	}

	return &entities.Transaction{
		Script:           t.Script,
		ReferenceBlockId: t.ReferenceBlockID[:],
		PayerAccount:     t.PayerAccount.Bytes(),
		ScriptAccounts:   scriptAccounts,
		Signatures:       signatures,
	}
}

func BlockHeaderToMessage(h *flow.Header) (entities.BlockHeader, error) {
	id := h.ID()
	bh := entities.BlockHeader{
		Id:       id[:],
		ParentId: h.ParentID[:],
		Height:   h.Height,
	}
	return bh, nil
}

func BlockToMessage(h *flow.Block) (*entities.Block, error) {

	parentID := h.ParentID
	t, err := ptypes.TimestampProto(h.Timestamp)
	if err != nil {
		return nil, err
	}

	cg := make([]*entities.CollectionGuarantee, len(h.Guarantees))
	for i, g := range h.Guarantees {
		cg[i] = collectionGuaranteeToMessage(g)
	}

	seals := make([]*entities.BlockSeal, len(h.Seals))
	for i, s := range h.Seals {
		seals[i] = blockSealToMessage(s)
	}

	bh := entities.Block{
		Height:               h.Height,
		ParentId:             parentID[:],
		Timestamp:            t,
		CollectionGuarantees: cg,
		BlockSeals:           seals,
		Signatures:           [][]byte{h.ParentVoterSig},
	}
	return &bh, nil
}

func collectionGuaranteeToMessage(g *flow.CollectionGuarantee) *entities.CollectionGuarantee {
	id := g.ID()

	return &entities.CollectionGuarantee{
		CollectionId: id[:],
		Signatures:   [][]byte{g.Signature},
	}
}

func blockSealToMessage(s *flow.Seal) *entities.BlockSeal {
	id := s.BlockID
	result := s.ExecutionResultID
	return &entities.BlockSeal{
		BlockId:                    id[:],
		ExecutionReceiptId:         result[:],
		ExecutionReceiptSignatures: [][]byte{s.Signature},
	}
}

func CollectionToMessage(c *flow.Collection) (*entities.Collection, error) {
	if c == nil || c.Transactions == nil {
		return nil, fmt.Errorf("invalid collection")
	}

	transactionsIDs := make([][]byte, len(c.Transactions))
	for i, t := range c.Transactions {
		id := t.ID()
		transactionsIDs[i] = id[:]
	}

	collectionID := c.ID()
	ce := &entities.Collection{
		Id:             collectionID[:],
		TransactionIds: transactionsIDs,
	}
	return ce, nil
}

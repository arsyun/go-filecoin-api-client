package iface

import (
	"context"
	"io"
)

type Asks struct {
	Miner  string      `json:"miner"`
	Price  string      `json:"price"`
	Expiry int64       `json:"expiry"`
	ID     int         `json:"id"`
	Error  interface{} `json:"error"`
}

type StorageDealPropose struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	DealID  string `json:"dealid"`
}

type StorageDealStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StorageDealInfo struct {
	State       string            `json:"state"`
	Message     string            `json:"message"`
	ProposalCid map[string]string `json:"proposal_cid"`
	ProofInfo   ProofInfo         `json:"proof_info"`
	Signature   string            `json:"signature"`
}

type ProofInfo struct {
	SectorID            uint64            `json:"sectorID"`
	CommD               []byte            `json:"commd"`
	CommR               []byte            `json:"commr"`
	CommRStar           []byte            `json:"comm_r_star"`
	CommitmentMessage   map[string]string `json:"commitment_message"`
	PieceInclusionProof []byte            `json:"piece_inclusion_proof"`
}

type StorageDealVerify struct {
}

type ClientAPI interface {
	Cat(ctx context.Context, cid string) (io.Reader, error)

	Import(ctx context.Context, fr io.Reader) (string, error)

	ListAsks(context.Context) ([]Asks, error)

	Payments(ctx context.Context, dealID string) ([]PaymentsDeal, error)

	//TODO: Storage order cannot be created at this time, pending verification
	ProposeStorageDeal(ctx context.Context, miner string, cid string, askId int, time int64) (StorageDealInfo, error)

	//TODO: pending verification
	QueryStorageDeal(ctx context.Context, dealID string) (StorageDealInfo, error)

	//TODO
	VerifyStorageDeal(ctx context.Context, dealID string) (StorageDealVerify, error)
}

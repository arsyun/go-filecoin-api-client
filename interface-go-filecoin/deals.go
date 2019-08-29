package iface

import (
	"context"
	fopt "go-filecoin-api-client/options"
)

type PaymentsDeal struct {
	Channel   int                  `json:"channel`
	Payer     string               `json:"payer"`
	Target    string               `json:"target"`
	Amount    string               `json:"amount"`
	ValidAt   int64                `json:"valid_at"`
	Condition PaymentDealCondition `json:"condition"`
	Signature string               `json:"signature"`
}

type PaymentDealCondition struct {
	To     string   `json:"to"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type DealsShowResult struct {
	DealCid         interface{}       `json:"deal_cid"`
	State           int               `json:"state"`
	MinerAddress    string            `json:"miner_address"`
	DurationBlocks  int64             `json:"duration_blocks"`
	DealSize        string            `json:"deal_size"`
	TotalPrice      string            `json:"total_price`
	PaymentVouchers []Paymentvouchers `json:"payment_vouchers"`
}

type Paymentvouchers struct {
	Index        int                  `json:"index"`
	Amount       string               `json:"amount"`
	ChannelID    string               `json:"channel_id"`
	Condition    PaymentDealCondition `json:"condition"`
	Payer        string               `json:"payer"`
	ValidAtBlock int64                `json:"valid_at_block"`
	EncodedAs    string               `json:"encode_as"`
}

type DealsRedeemResult struct {
	Cid     map[string]string `json:"cid"`
	GasUsed string            `json:"gas_used"`
	Preview bool              `json:"preview"`
}

type DealsInfo struct {
	MinerAddress string            `json:"miner_address"`
	PieceCid     map[string]string `json:"piece_cid"`
	ProposalCid  map[string]string `json:"proposal_cid"`
	State        string            `json:"state"`
}

type DealsAPI interface {
	List(context.Context) ([]DealsInfo, error)

	Redeem(ctx context.Context, dealID string, opts ...fopt.DealsOption) (DealsRedeemResult, error)

	Show(ctx context.Context, dealID string) (DealsShowResult, error)
}

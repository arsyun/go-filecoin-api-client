package iface

import (
	"context"
	fopt "go-filecoin-api-client/options"
)

type MinerPower struct {
	Power string `json:"power"`
	Total string `json:"total"`
}

type MinerProvingPeriod struct {
	Start      int64                  `json:"start"`
	End        int64                  `json:"end"`
	ProvingSet map[string]interface{} `json:"proving_set"`
}

type MinerCreateResult struct {
	Address map[string]string `json:"address"`
	GasUsed string            `json:"gas_used"`
	Preview bool              `json:"preview"`
}

type MinerSetPriceResult struct {
	GasUsed               string                `json:"gas_used"`
	MinerSetPriceResponse MinerSetPriceResponse `json:"miner_set_price_response"`
	Preview               bool                  `json:"preview"`
}

type MinerUpdatePeerIDResult struct {
	Cid     map[string]string `json:"cid"`
	GasUsed string            `json:"gas_used"`
	Preview bool              `json:"preview"`
}

type MinerSetPriceResponse struct {
	AddAskCid map[string]string `json:"add_ask_cid"`
	BlockCid  map[string]string `json:"block_cid"`
	MinerAddr string            `json:"miner_addr"`
	Price     string            `json:"price"`
}

type MinerAPI interface {
	Collateral(ctx context.Context, minerAddr string) (string, error)

	Create(ctx context.Context, collateral int64, opts ...fopt.MinerCreateOption) (MinerCreateResult, error)

	Owner(ctx context.Context, minerAddr string) (actorAddr string, err error)

	Power(ctx context.Context, minerAddr string) (MinerPower, error)

	ProvingPeriod(ctx context.Context, minerAddr string) (MinerProvingPeriod, error)

	SetPrice(ctx context.Context, storageprice int64, expiry int64, opts ...fopt.MinerSetPriceOption) (MinerSetPriceResult, error)

	UpdatePeerId(ctx context.Context, addr string, peerid string, opts ...fopt.MinerUpdatePeerIdOption) (MinerUpdatePeerIDResult, error)
}

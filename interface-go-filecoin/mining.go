package iface

import (
	"context"
)

type MiningCofig struct {
	MinerAddress            string `json:"miner_address"`
	AutoSealIntervalSeconds int64  `json:"auto_seal_interval_seconds"`
	StoragePrice            string `json:"storage_price"`
}

type MiningStatusResult struct {
	Active        bool               `json:"active"`
	MinerAddress  string             `json:"miner_address"`
	Owner         string             `json:"owner"`
	Collateral    string             `json:"collateral"`
	ProvingPeriod MinerProvingPeriod `json:"proving_period"`
}

type MiningAPI interface {
	Address(ctx context.Context) (minerAddr string, err error)
	//TODO
	Once(ctx context.Context) (BlockMsg, error)

	SealNow(ctx context.Context) (error)

	Start(ctx context.Context) (error)

	Status(ctx context.Context) (MiningStatusResult, error)

	Stop(ctx context.Context) (error)
}

package iface

import (
	"context"
)

type BitSwapStats struct {
	Providebuflen   int      `json:"providebuflen"`
	Wantlist        []string `json:"wantlist"`
	Peers           []string `json:"peers"`
	BlocksReceived  int      `json:"blocks_received"`
	DataReceived    int      `json:"data_received"`
	BlocksSent      int      `json:"blocks_sent"`
	DataSent        int      `json:"data_sent"`
	DupBlksReceived int      `json:"dup_blks_received"`
	MessageReceived int64    `json:"message_received"`
}

type BitSwapAPI interface {
	Stats(context.Context) (BitSwapStats, error)
}

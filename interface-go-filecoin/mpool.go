package iface

import (
	"context"
)

type MpoolConfig struct {
	MaxPoolSize int64  `json:"max_pool_size"`
	MaxNonceGap string `json:"max_nonce_gap"`
}

type MpoolAPI interface {
	Ls(context.Context) ([]MessageDetail, error)

	Rm(ctx context.Context, msgid string) (error)

	Show(ctx context.Context, msgid string) (MessageDetail, error)
}

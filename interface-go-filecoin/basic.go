package iface

import (
	"context"
)

type IDResult struct {
	Addresses []string `json:"addresses"`
	ID        string   `json:"id"`
}

type VersionResult struct {
	Commit string `json:commit`
}

type BasicAPI interface {
	Daemon(ctx context.Context) (error)

	ID(ctx context.Context) (IDResult, error)

	//TODO
	Ping(ctx context.Context, peerID string) (string, error)

	Version(ctx context.Context) (VersionResult, error)
}

package iface

import (
	"context"
	fopt "go-filecoin-api-client/options"
)

type LogAPI interface {
	Level(ctx context.Context, level string, opts ...fopt.LogLevelOption) (string, error)

	Ls(context.Context) ([]string, error)
	//TODO
	Tail(context.Context) (string, error)
}

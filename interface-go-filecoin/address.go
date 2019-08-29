package iface

import (
	"context"
)

type AddressAPI interface {
	Default(context.Context) (string, error)

	Ls(context.Context) ([]string, error)

	Lookup(ctx context.Context, minerAddr string) (string, error)

	New(context.Context) (string, error)
}

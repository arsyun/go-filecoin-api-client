package iface

import (
	"context"
)

type Chain struct {
	l string
}

type ChainAPI interface {
	Head(context.Context) ([]string, error)

	Ls(context.Context) (string, error)
}

package iface

import (
	"context"
)

type DagGetInfo struct {
	Data  string   `json:"data"`
	Links []string `json:"links"`
}

type DagAPI interface {
	Get(ctx context.Context, cid string) (DagGetInfo, error)
}

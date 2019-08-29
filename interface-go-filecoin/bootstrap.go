package iface

import (
	"context"
)

type BootstrapConfig struct {
	Addresses        []string `json:"address"`
	MinPeerThreshold int      `json:"min_peer_threshold"`
	Period           string   `json:"period"`
}

type BootStrapAPI interface {
	Ls(context.Context) ([]string, error)
}

package iface

import (
	"context"
)

type SwarmConfig struct {
	Address string `json:"address"`
}

type SwarmPeers struct {
	Peers []PeerInfo `json:"peers"`
}

type PeerInfo struct {
	Addr    string      `json:"addr"`
	Peer    string      `json:"peer"`
	Latency string      `json:"latency"`
	Muxer   string      `json:"muxer"`
	Streams interface{} `json:"streams"`
}

type SwarmAPI interface {
	Connect(ctx context.Context, addrs ...string) (connectIds []string, err error)

	Peers(ctx context.Context) (SwarmPeers, error)
}

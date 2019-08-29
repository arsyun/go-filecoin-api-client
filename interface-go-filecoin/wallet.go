package iface

import (
	"context"
)

type WalletConfig struct {
	DefaultAddress string `json:"default_address"`
}

type WalletInfo struct {
	KeyInfo []KeyInfo `json:"key_info"`
}

type KeyInfo struct {
	Privatekey string `json:"privatekey"`
	Curve      string `json:"curve"`
}

type WalletImportResult struct {
	Addresses []string `json:"addresses"`
}

type WalletAPI interface {
	Balance(ctx context.Context, address string) (string, error)

	Export(ctx context.Context, addrs ...string) (WalletInfo, error)

	Import(ctx context.Context, keyinfos []KeyInfo) ([]string, error)
}

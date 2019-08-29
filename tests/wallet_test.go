package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var WalletAPI = api.NewAPI("").Wallet()

func TestWalletBalance(t *testing.T) {
	addr, err := AddrAPI.Default(context.Background())
	require.NoError(t, err)

	resp, err := WalletAPI.Balance(context.Background(), addr)
	require.NoError(t, err)
	t.Log("Wallet balance:", resp)
}

func TestWalletExport(t *testing.T) {
	addr, err := AddrAPI.Default(context.Background())
	require.NoError(t, err)

	resp, err := WalletAPI.Export(context.Background(), addr)
	require.NoError(t, err)
	t.Log("Wallet export:", resp)
}

func TestWalletImport(t *testing.T) {
	addr, err := AddrAPI.Default(context.Background())
	require.NoError(t, err)

	WalletInfo, err := WalletAPI.Export(context.Background(), addr)
	require.NoError(t, err)

	resp, err := WalletAPI.Import(context.Background(), WalletInfo.KeyInfo)
	require.NoError(t, err)
	t.Log("Wallet Import:", resp)
}

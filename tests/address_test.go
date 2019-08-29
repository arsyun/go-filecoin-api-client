package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var AddrAPI = api.NewAPI("").Address()

func TestAddressDefault(t *testing.T) {
	resp, err := AddrAPI.Default(context.Background())
	require.NoError(t, err)
	t.Log("default address:", resp)
}

func TestAddressLs(t *testing.T) {
	resp, err := AddrAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("address list:", resp)
}

func TestAddressLookup(t *testing.T) {
	addrs, err := AddrAPI.Ls(context.Background())
	require.NoError(t, err)

	resp, err := AddrAPI.Lookup(context.Background(), addrs[0])
	require.NoError(t, err)
	t.Log("peerid of address:", resp)
}

func TestAddressNew(t *testing.T) {
	resp, err := AddrAPI.New(context.Background())
	require.NoError(t, err)
	t.Log("new address:", resp)
}

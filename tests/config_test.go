package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var ConfigAPI = api.NewAPI("").Config()

func TestConfigGet(t *testing.T) {
	resp, err := ConfigAPI.Get(context.Background(), "wallet.defaultAddress")
	require.NoError(t, err)
	t.Log("get config value:", resp)
}

func TestConfigSet(t *testing.T) {
	resp, err := ConfigAPI.Set(context.Background(), "mpool.maxNonceGap", "101")
	require.NoError(t, err)
	t.Log("set config value:", resp)
}

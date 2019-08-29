package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var BitSwapAPI = api.NewAPI("").BitSwap()

func TestBitSwapStats(t *testing.T) {
	resp, err := BitSwapAPI.Stats(context.Background())
	require.NoError(t, err)
	t.Log("bitswap statistics:", resp)
}

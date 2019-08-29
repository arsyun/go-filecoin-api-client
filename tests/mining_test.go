package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var MiningAPI = api.NewAPI("").Mining()

func TestMiningAddress(t *testing.T) {
	resp, err := MiningAPI.Address(context.Background())
	require.NoError(t, err)
	t.Log("miner address:", resp)
}

func TestMiningSealNow(t *testing.T) {
	err := MiningAPI.SealNow(context.Background())
	require.NoError(t, err)
	t.Log("Start sealing success")
}

func TestMiningStart(t *testing.T) {
	err := MiningAPI.Start(context.Background())
	require.NoError(t, err)

	minerStatus, err := MiningAPI.Status(context.Background())
	require.NoError(t, err)

	if minerStatus.Active {
		t.Log("Node is already mining")
	} else {
		t.Log("Started mining")
	}
}

func TestMiningStatus(t *testing.T) {
	resp, err := MiningAPI.Status(context.Background())
	require.NoError(t, err)
	t.Log("Mining Status", resp)
}

func TestMiningStop(t *testing.T) {
	err := MiningAPI.Stop(context.Background())
	require.NoError(t, err)
	t.Log("Stopped mining")
}

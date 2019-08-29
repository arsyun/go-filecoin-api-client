package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var SwarmAPI = api.NewAPI("").Swarm()

func TestSwarmPeers(t *testing.T) {
	resp, err := SwarmAPI.Peers(context.Background())
	require.NoError(t, err)
	t.Log("List peers with open connections:", resp)
}

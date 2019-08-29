package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	fopt "go-filecoin-api-client/options"
	"testing"
)

var DhtAPI = api.NewAPI("").DHT()

func TestDhtFindPeer(t *testing.T) {
	resp, err := DhtAPI.FindPeer(context.Background(), "peerid")
	require.NoError(t, err)
	t.Log("peerid of multiaddresses:", resp)
}

func TestDhtFindProvs(t *testing.T) {
	resp, err := DhtAPI.FindProvs(context.Background(), "id", fopt.DHT.NumProviders(1))
	require.NoError(t, err)
	t.Log("peers:", resp)
}

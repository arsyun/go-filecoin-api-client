package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	fopt "go-filecoin-api-client/options"
	"testing"
)

var LogAPI = api.NewAPI("").Log()

func TestLogLevel(t *testing.T) {
	resp, err := LogAPI.Level(context.Background(), "error", fopt.Log.Subsystem("ping"))
	require.NoError(t, err)
	t.Log("peerid of multiaddresses:", resp)
}

func TestLogLs(t *testing.T) {
	resp, err := LogAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("peers:", resp)
}

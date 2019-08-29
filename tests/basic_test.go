package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var BasicAPI = api.NewAPI("").Basic()

func TestDaemon(t *testing.T) {
	err := BasicAPI.Daemon(context.Background())
	require.NoError(t, err)
	t.Log("daemon success")
}

func TestID(t *testing.T) {
	resp, err := BasicAPI.ID(context.Background())
	require.NoError(t, err)
	t.Log("network peers info:", resp)
}

func TestVersion(t *testing.T) {
	resp, err := BasicAPI.Version(context.Background())
	require.NoError(t, err)
	t.Log("go-filecoin version information:", resp)
}

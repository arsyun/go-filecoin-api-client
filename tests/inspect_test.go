package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var InspectAPI = api.NewAPI("").Inspect()

func TestInspectAll(t *testing.T) {
	resp, err := InspectAPI.All(context.Background())
	require.NoError(t, err)
	t.Log("all diagnostic info:", resp)
}

func TestInspectConfig(t *testing.T) {
	resp, err := InspectAPI.Config(context.Background())
	require.NoError(t, err)
	t.Log("in-memory config info:", resp)
}

func TestInspectDisk(t *testing.T) {
	resp, err := InspectAPI.Disk(context.Background())
	require.NoError(t, err)
	t.Log("filesystem usage info:", resp)
}

func TestInspectEnvironment(t *testing.T) {
	resp, err := InspectAPI.Environment(context.Background())
	require.NoError(t, err)
	t.Log("filecoin environment info:", resp)
}

func TestInspectMemory(t *testing.T) {
	resp, err := InspectAPI.Memory(context.Background())
	require.NoError(t, err)
	t.Log("memory usage info:", resp)
}

func TestInspectRuntime(t *testing.T) {
	resp, err := InspectAPI.Runtime(context.Background())
	require.NoError(t, err)
	t.Log("runtime diagnostic info:", resp)
}

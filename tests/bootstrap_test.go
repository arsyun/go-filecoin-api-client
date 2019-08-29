package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var BootStrapAPI = api.NewAPI("").BootStrap()

func TestBootStrapLs(t *testing.T) {
	resp, err := BootStrapAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("bootstrap list:", resp)
}

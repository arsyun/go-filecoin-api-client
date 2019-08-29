package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var ChainAPI = api.NewAPI("").Chain()

func TestChainHead(t *testing.T) {
	resp, err := ChainAPI.Head(context.Background())
	require.NoError(t, err)
	t.Log("heaviest tipset CIDs:", resp)
}

func TestChainLs(t *testing.T) {
	resp, err := ChainAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("blocks list:", resp)
}

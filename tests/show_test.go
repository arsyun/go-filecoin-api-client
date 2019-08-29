package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var ShowAPI = api.NewAPI("").Show()

func TestShowBlock(t *testing.T) {
	headers, err := ChainAPI.Head(context.Background())
	require.NoError(t, err)

	resp, err := ShowAPI.Block(context.Background(), headers[0])
	require.NoError(t, err)
	t.Log("Show a full filecoin block:", resp)
}

func TestShowHeader(t *testing.T) {
	headers, err := ChainAPI.Head(context.Background())
	require.NoError(t, err)

	resp, err := ShowAPI.Header(context.Background(), headers[0])
	require.NoError(t, err)
	t.Log("Show a filecoin message collection:", resp)
}

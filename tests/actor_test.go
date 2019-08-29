package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var ActorAPI = api.NewAPI("").Actor()

func TestActorLs(t *testing.T) {
	resp, err := ActorAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("actor list:", resp)
}

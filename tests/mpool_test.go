package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var MpoolAPI = api.NewAPI("").Mpool()

func TestMpoolLs(t *testing.T) {
	resp, err := MpoolAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("message list:", resp)
}

func TestMpoolRm(t *testing.T) {
	err := MpoolAPI.Rm(context.Background(), "msgid")
	require.NoError(t, err)
	t.Log("delete message success")
}

func TestMpoolShow(t *testing.T) {
	resp, err := MpoolAPI.Show(context.Background(), "msgid")
	require.NoError(t, err)
	t.Log("message info:", resp)
}

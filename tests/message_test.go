package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	fopt "go-filecoin-api-client/options"
	"testing"
)

var MessageAPI = api.NewAPI("").Message()

func TestMessageSend(t *testing.T) {
	resp, err := MessageAPI.Send(context.Background(), "", "", fopt.Message.GasLimit("0"), fopt.Message.GasPrice("0.1"))
	require.NoError(t, err)
	t.Log("send a message:", resp)
}

func TestMessageStatus(t *testing.T) {
	msgCid, err := MessageAPI.Send(context.Background(), "", "", fopt.Message.GasLimit("0"), fopt.Message.GasPrice("0.1"))
	require.NoError(t, err)

	resp, err := MessageAPI.Status(context.Background(), msgCid)
	require.NoError(t, err)
	t.Log("message status:", resp)
}

func TestMessageWait(t *testing.T) {
	msgCid, err := MessageAPI.Send(context.Background(), "", "", fopt.Message.GasLimit("0"), fopt.Message.GasPrice("1"))
	require.NoError(t, err)

	resp, err := MessageAPI.Wait(context.Background(), msgCid)
	require.NoError(t, err)
	t.Log("wait a message:", resp)
}

package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	fopt "go-filecoin-api-client/options"
	"testing"
)

var PaychAPI = api.NewAPI("").Paych()

func TestPaychCancel(t *testing.T) {
	resp, err := PaychAPI.Cancel(context.Background(), 455, fopt.Paych.GasPrice("0.1"), fopt.Paych.GasLimit("100"))
	require.NoError(t, err)
	t.Log("Cancel a payment channel:", resp)
}

func TestPaychClose(t *testing.T) {
	voucher, err := PaychAPI.Voucher(context.Background(), 14, "0.0000001")
	require.NoError(t, err)

	resp, err := PaychAPI.Close(context.Background(), voucher, fopt.Paych.GasPrice("0.1"), fopt.Paych.GasLimit("100"))
	require.NoError(t, err)
	t.Log("Close a payment channel:", resp)
}

func TestPaychCreate(t *testing.T) {
	resp, err := PaychAPI.Create(context.Background(), "t1627jvaeugx5bhpvt6tvb4jfhttohikg6kmljnka", "0.0001", 54, fopt.Paych.GasPrice("0.1"), fopt.Paych.GasLimit("100"))
	require.NoError(t, err)
	t.Log("Create new payment channel:", resp)
}

func TestPaychLs(t *testing.T) {
	resp, err := PaychAPI.Ls(context.Background())
	require.NoError(t, err)
	t.Log("Payment channels list:", resp)
}

func TestPaychExtend(t *testing.T) {
	resp, err := PaychAPI.Extend(context.Background(), 755, "0.000001", 145, fopt.Paych.GasPrice("0.1"), fopt.Paych.GasLimit("100"))
	require.NoError(t, err)
	t.Log("Extend the value and lifetime of the channel:", resp)
}

func TestPaychReclaim(t *testing.T) {
	resp, err := PaychAPI.Reclaim(context.Background(), 957, fopt.Paych.GasPrice("0.1"), fopt.Paych.GasLimit("100"))
	require.NoError(t, err)
	t.Log("Reclaim funds from an expired channel:", resp)
}

func TestPaychRedeem(t *testing.T) {
	voucher, err := PaychAPI.Voucher(context.Background(), 14, "0.000001")
	require.NoError(t, err)

	resp, err := PaychAPI.Redeem(context.Background(), voucher)
	require.NoError(t, err)
	t.Log("Redeem a payment voucher against a payment channel:", resp)
}

func TestPaychVoucher(t *testing.T) {
	resp, err := PaychAPI.Voucher(context.Background(), 14, "0.0001")
	require.NoError(t, err)
	t.Log("Create a new voucher from a payment channel:", resp)
}

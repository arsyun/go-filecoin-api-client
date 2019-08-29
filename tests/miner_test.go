package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	fopt "go-filecoin-api-client/options"
	"testing"
)

var MinerAPI = api.NewAPI("").Miner()

func TestMinerCollateral(t *testing.T) {
	resp, err := MinerAPI.Collateral(context.Background(), "t2sbfa6bellsm3vgod3v3j42cadmcuigzzyc5lcyy")
	require.NoError(t, err)
	t.Log("miner collateral:", resp)
}

//TODO
func TestMinerCreate(t *testing.T) {
	//resp, err := MinerAPI.Create(context.Background(), 500, fopt.Miner.CreateMinerGasPrice("0.001"), fopt.Miner.CreateMinerGasLimit("100"))
	//require.NoError(t, err)
	//t.Log("create miner:", resp)
}

func TestMinerOwner(t *testing.T) {
	resp, err := MinerAPI.Owner(context.Background(), "t2sbfa6bellsm3vgod3v3j42cadmcuigzzyc5lcyy")
	require.NoError(t, err)
	t.Log("actor address of miner:", resp)
}

func TestMinerPower(t *testing.T) {
	resp, err := MinerAPI.Power(context.Background(), "t2sbfa6bellsm3vgod3v3j42cadmcuigzzyc5lcyy")
	require.NoError(t, err)
	t.Log("the power of a miner and the total storage market: ", resp)
}

func TestMinerProvingPeriod(t *testing.T) {
	resp, err := MinerAPI.ProvingPeriod(context.Background(), "t2sbfa6bellsm3vgod3v3j42cadmcuigzzyc5lcyy")
	require.NoError(t, err)
	t.Log("Proving Period: ", resp)
}

func TestMinerSetPrice(t *testing.T) {
	resp, err := MinerAPI.SetPrice(context.Background(), 100, 2800, fopt.Miner.SetPriceGasPrice("0.1"), fopt.Miner.SetPriceGasLimit("100"))
	require.NoError(t, err)
	t.Log("Set the minimum price for storage: ", resp)
}

func TestMinerUpdatePeerId(t *testing.T) {
	resp, err := MinerAPI.UpdatePeerId(context.Background(), "addr", "peerid", fopt.Miner.UpPeerIdGasPrice("0.1"), fopt.Miner.UpPeerIdGasLimit("100"))
	require.NoError(t, err)
	t.Log("Change the libp2p identity that a miner is operating: ", resp)
}

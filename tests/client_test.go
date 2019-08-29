package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"io/ioutil"
	"strings"
	"testing"
)

var ClientAPI = api.NewAPI("").Client()

func TestClientCat(t *testing.T) {
	fr := strings.NewReader("test client import")
	cid, err := ClientAPI.Import(context.Background(), fr)
	require.NoError(t, err)

	resp, err := ClientAPI.Cat(context.Background(), cid)
	require.NoError(t, err)

	info, err := ioutil.ReadAll(resp)
	t.Log("Read out data stored on the network:", string(info))
}

func TestClientImport(t *testing.T) {
	fr := strings.NewReader("test client import")
	resp, err := ClientAPI.Import(context.Background(), fr)
	require.NoError(t, err)
	t.Log("File cid:", resp)
}

func TestClientListAsks(t *testing.T) {
	resp, err := ClientAPI.ListAsks(context.Background())
	require.NoError(t, err)
	t.Log("All asks in the storage market:", resp)
}

func TestClientPayments(t *testing.T) {
	resp, err := ClientAPI.Payments(context.Background(), "dealidxxx")
	require.NoError(t, err)
	t.Log("Payments for a given deal:", resp)
}

//TODO
func TestClientProposeStorageDeal(t *testing.T) {
	//miner := "t2crutvsqw7m4p7ljf4gfyaf4jr6zxymnhisuw4oi"
	//cid := "QmQXLZgtEXu2dADSqA8rhEpnw7vzpJLATVEKDQKKc9Yhiw"
	//askid := 0
	//time := int64(2800) //one day
	//resp, err := ClientAPI.ProposeStorageDeal(context.Background(), miner, cid, askid, time)
	//require.NoError(t, err)
	//t.Log("Sends a storage deal proposal to a miner:", resp)
}

func TestClientQueryStorageDeal(t *testing.T) {
	resp, err := ClientAPI.QueryStorageDeal(context.Background(), "dealId")
	require.NoError(t, err)
	t.Log("deal's status:", resp)
}

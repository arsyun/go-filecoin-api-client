package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"strings"
	"testing"
)

var DagAPI = api.NewAPI("").Dag()

func TestDagGet(t *testing.T) {

	fr := strings.NewReader("test dag get")
	cid, err := ClientAPI.Import(context.Background(), fr)
	require.NoError(t, err)
	t.Log("import file cid:", cid)

	resp, err := DagAPI.Get(context.Background(), cid)
	require.NoError(t, err)
	t.Log("dag node:", resp)
}

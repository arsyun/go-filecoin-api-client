package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-filecoin-api-client/api"
	"testing"
)

var StatsAPI = api.NewAPI("").Stats()

func TestStatsBandwidth(t *testing.T) {
	resp, err := StatsAPI.BandWidth(context.Background())
	require.NoError(t, err)
	t.Log("View bandwidth usage metrics:", resp)
}

package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type BitSwapAPI HttpApi

func (api *BitSwapAPI) Stats(ctx context.Context) (iface.BitSwapStats, error) {
	var out iface.BitSwapStats
	err := api.core().Request("bitswap/stats").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *BitSwapAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

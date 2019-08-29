package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type StatsAPI HttpApi

func (api *StatsAPI) BandWidth(ctx context.Context) (iface.StatsBandWidth, error) {
	var out iface.StatsBandWidth

	err := api.core().Request("stats/bandwidth").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *StatsAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

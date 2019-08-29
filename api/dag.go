package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type DagAPI HttpApi

func (api *DagAPI) Get(ctx context.Context, cid string) (iface.DagGetInfo, error) {
	var out iface.DagGetInfo

	err := api.core().Request("dag/get", cid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *DagAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

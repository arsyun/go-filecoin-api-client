package api

import (
	"context"
)

type BootStrapAPI HttpApi

func (api *BootStrapAPI) Ls(ctx context.Context)([]string, error){
	var out struct{
		Peers []string
	}

	err := api.core().Request("bootstrap/ls").
		Exec(ctx, &out)
	if err != nil {
		return out.Peers, err
	}

	return out.Peers, nil
}

func (api *BootStrapAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

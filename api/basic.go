package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type BasicAPI HttpApi

func (api *BasicAPI) Daemon(ctx context.Context) (error) {
	err := api.core().Request("daemon").
		Exec(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *BasicAPI) ID(ctx context.Context) (iface.IDResult, error) {
	var out iface.IDResult
	err := api.core().Request("id").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *BasicAPI) Version(ctx context.Context) (iface.VersionResult, error) {
	var out iface.VersionResult
	err := api.core().Request("version").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

//TODO
func (api *BasicAPI) Ping(ctx context.Context, peerID string) (string, error) {
	return "", nil
}

func (api *BasicAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

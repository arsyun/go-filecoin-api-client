package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type InspectAPI HttpApi

func (api *InspectAPI) All(ctx context.Context) (iface.InspectAll, error) {
	var out iface.InspectAll
	err := api.core().Request("inspect/all").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *InspectAPI) Config(ctx context.Context) (iface.Config, error) {
	var out iface.Config
	err := api.core().Request("inspect/config").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *InspectAPI) Disk(ctx context.Context) (iface.Disk, error) {
	var out iface.Disk
	err := api.core().Request("inspect/disk").
		Exec(ctx, &out)
	if err != nil {
		return out, nil
	}
	return out, nil
}

func (api *InspectAPI) Environment(ctx context.Context) (iface.Environment, error) {
	var out iface.Environment
	err := api.core().Request("inspect/environment").
		Exec(ctx, &out)
	if err != nil {
		return out, nil
	}
	return out, nil
}

func (api *InspectAPI) Memory(ctx context.Context) (iface.Memory, error) {
	var out iface.Memory
	err := api.core().Request("inspect/memory").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *InspectAPI) Runtime(ctx context.Context) (iface.Runtime, error) {
	var out iface.Runtime
	err := api.core().Request("inspect/runtime").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *InspectAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

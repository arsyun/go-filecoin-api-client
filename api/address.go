package api

import (
	"context"
)

type AddressAPI HttpApi

func (api *AddressAPI) Default(ctx context.Context) (string, error) {
	var out struct {
		Address string
	}
	err := api.core().Request("address/default").
		Exec(ctx, &out)
	if err != nil {
		return "", err
	}
	return out.Address, nil
}

func (api *AddressAPI) Ls(ctx context.Context) ([]string, error) {
	var out struct {
		Addresses []string
	}
	err := api.core().Request("address/ls").
		Exec(ctx, &out)
	if err != nil {
		return nil, err
	}
	return out.Addresses, nil
}

func (api *AddressAPI) Lookup(ctx context.Context, minerAddr string) (string, error) {
	var out string
	err := api.core().Request("address/lookup", minerAddr).
		Exec(ctx, &out)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (api *AddressAPI) New(ctx context.Context) (string, error) {
	var out struct {
		Address string
	}
	err := api.core().Request("address/new").
		Exec(ctx, &out)
	if err != nil {
		return "", err
	}
	return out.Address, nil
}

func (api *AddressAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

package api

import(
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type MiningAPI HttpApi

func (api *MiningAPI) Address(ctx context.Context) (minerAddr string, err error) {
	err = api.core().Request("mining/address").
		Exec(ctx, &minerAddr)
	if err != nil {
		return "", err
	}

	return minerAddr, nil
}

func (api *MiningAPI) Once(ctx context.Context) (iface.BlockMsg, error) {
	var out iface.BlockMsg
	return out, nil
}

func (api *MiningAPI) SealNow(ctx context.Context) (error) {
	err := api.core().Request("mining/seal-now").
		Exec(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (api *MiningAPI) Start(ctx context.Context) (error) {
	err := api.core().Request("mining/start").
		Exec(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *MiningAPI) Status(ctx context.Context) (iface.MiningStatusResult, error) {
	var out iface.MiningStatusResult

	err := api.core().Request("mining/status").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MiningAPI) Stop(ctx context.Context) (error) {
	err := api.core().Request("mining/stop").
		Exec(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (api *MiningAPI) core() *HttpApi {
	return (*HttpApi)(api)
}
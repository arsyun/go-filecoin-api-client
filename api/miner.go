package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
	fopt "go-filecoin-api-client/options"
	"strconv"
)

type MinerAPI HttpApi

func (api *MinerAPI) Collateral(ctx context.Context, minerAddr string) (collateral string, err error) {
	err = api.core().Request("miner/collateral", minerAddr).
		Exec(ctx, &collateral)
	if err != nil {
		return "", err
	}
	return collateral, nil
}

func (api *MinerAPI) Create(ctx context.Context, collateral int64, opts ...fopt.MinerCreateOption) (iface.MinerCreateResult, error) {
	var out iface.MinerCreateResult

	options, err := fopt.MinerCreateOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("miner/create", strconv.FormatInt(collateral, 10)).
		Option("gas-price", options.GasPrice).
		Option("gas-limit", options.GasLimit)
	// defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if options.PeerId != "" {
		req = req.Option("peerid", options.PeerId)
	}
	if options.Sectorsize != 0 {
		req = req.Option("sectorsize", options.Sectorsize)
	}

	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (api *MinerAPI) Owner(ctx context.Context, minerAddr string) (actorAddr string, err error) {

	err = api.core().Request("miner/owner", minerAddr).
		Exec(ctx, &actorAddr)
	if err != nil {
		return "", err
	}
	return actorAddr, nil
}

func (api *MinerAPI) Power(ctx context.Context, minerAddr string) (iface.MinerPower, error) {
	var out iface.MinerPower

	err := api.core().Request("miner/power", minerAddr).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *MinerAPI) ProvingPeriod(ctx context.Context, minerAddr string) (iface.MinerProvingPeriod, error) {
	var out iface.MinerProvingPeriod

	err := api.core().Request("miner/proving-period", minerAddr).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (api *MinerAPI) SetPrice(ctx context.Context, storageprice int64, expiry int64, opts ...fopt.MinerSetPriceOption) (iface.MinerSetPriceResult, error) {
	var out iface.MinerSetPriceResult

	options, err := fopt.MinerSetPriceOptions(opts...)
	if err != nil {
		return out, err
	}

	err = api.core().Request("miner/set-price").
		Option("gas-price", options.GasPrice).
		Option("gas-limit", options.GasLimit).
		Arguments(strconv.FormatInt(storageprice, 10)).
		Arguments(strconv.FormatInt(expiry, 10)).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MinerAPI) UpdatePeerId(ctx context.Context, addr string, peerId string, opts ...fopt.MinerUpdatePeerIdOption) (iface.MinerUpdatePeerIDResult, error) {
	var out iface.MinerUpdatePeerIDResult

	options, err := fopt.MinerUpdatePeerIdOptions(opts...)
	if err != nil {
		return out, err
	}

	err = api.core().Request("miner/update-peerid", addr, peerId).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MinerAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

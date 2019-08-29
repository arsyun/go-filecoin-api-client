package api

import (
	"bytes"
	"context"
	"encoding/json"
	iface "go-filecoin-api-client/interface-go-filecoin"
	fopt "go-filecoin-api-client/options"
	"io"
	"strings"
)

type DealsAPI HttpApi

func (api *DealsAPI) List(ctx context.Context) ([]iface.DealsInfo, error) {
	var out []iface.DealsInfo

	resp, err := api.core().Request("deals/list").
		Send(ctx)
	if err != nil {
		return out, err
	}
	if resp.Error != nil {
		return out, resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err = io.Copy(b, resp.Output); err != nil {
		return out, err
	}

	a := "[" + strings.TrimRight(strings.ReplaceAll(b.String(), "\n", ","), ",") + "]"
	if err = json.Unmarshal([]byte(a), &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *DealsAPI) Redeem(ctx context.Context, dealID string, opts ...fopt.DealsOption) (iface.DealsRedeemResult, error) {
	var out iface.DealsRedeemResult

	options, err := fopt.DealsOptions(opts...)
	if err != nil {
		return out, err
	}

	err = api.core().Request("deals/redeem").
		Option("gas-price", options.GasPrice).
		Option("gas-limit", options.GasLimit).
		Arguments(dealID).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *DealsAPI) Show(ctx context.Context, dealID string) (iface.DealsShowResult, error) {
	var out iface.DealsShowResult

	err := api.core().Request("deals/show", dealID).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *DealsAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

package api

import (
	"bytes"
	"context"
	"encoding/json"
	iface "go-filecoin-api-client/interface-go-filecoin"
	fopt "go-filecoin-api-client/options"
	"io"
	"strconv"
)

type PaychAPI HttpApi

func (api *PaychAPI) Cancel(ctx context.Context, channelID int64, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo

	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/cancel", strconv.FormatInt(channelID, 10)).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview)
	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}

	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Close(ctx context.Context, voucher string, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo
	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/close", voucher).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview)

	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Create(ctx context.Context, target string, amount string, eol int64, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo

	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/create").
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview).
		Arguments(target).
		Arguments(amount).
		Arguments(strconv.FormatInt(eol, 10))
	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}

	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Ls(ctx context.Context) (map[string]iface.Paych, error) {
	p := make(map[string]iface.Paych)
	resp, err := api.core().Request("paych/ls").
		Send(ctx)
	if err != nil {
		return nil, nil
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err = io.Copy(b, resp.Output); err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(b.String()), &p); err != nil {
		return nil, err
	}

	return p, nil
}

func (api *PaychAPI) Extend(ctx context.Context, channelID int64, amount string, eol int64, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo

	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/extend").
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview).
		Arguments(strconv.FormatInt(channelID, 10)).
		Arguments(amount).
		Arguments(strconv.FormatInt(eol, 10))

	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Reclaim(ctx context.Context, channelID int64, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo
	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/reclaim", strconv.FormatInt(channelID, 10)).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview)

	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Redeem(ctx context.Context, voucher string, opts ...fopt.PaychOption) (iface.ChannelInfo, error) {
	var out iface.ChannelInfo
	options, err := fopt.PaychOptions(opts...)
	if err != nil {
		return out, err
	}

	req := api.core().Request("paych/redeem", voucher).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("preview", options.Preview)

	//defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if err = req.Exec(ctx, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *PaychAPI) Voucher(ctx context.Context, channelID int64, amount string, opts ...fopt.PaychVoucherOption) (voucher string, err error) {
	options, err := fopt.PaychVoucherOptions()
	if err != nil {
		return "", err
	}

	req := api.core().Request("paych/voucher").
		Arguments(strconv.FormatInt(channelID, 10)).
		Arguments(amount)

	if options.Validat != 0 {
		req.Option("validat", options.Validat)
	}

	if err = req.Exec(ctx, &voucher); err != nil {
		return "", err
	}

	return voucher, nil
}

func (api *PaychAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

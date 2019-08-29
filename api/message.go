package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
	fopt "go-filecoin-api-client/options"
)

type MessageAPI HttpApi

func (api *MessageAPI) Send(ctx context.Context, target string, method string, opts ...fopt.MessageSendOption) (msgCid string, err error) {
	var out iface.MsgSendResult

	options, err := fopt.MessageSendOptions(opts...)
	if err != nil {
		return "", err
	}

	req := api.core().Request("message/send", target, method).
		Option("gas-limit", options.GasLimit).
		Option("gas-price", options.GasPrice).
		Option("value", options.Value)

	//from is defaultAddress
	if options.From != "" {
		req = req.Option("from", options.From)
	}
	if err = req.Exec(ctx, &out); err != nil {
		return "", err
	}

	if _, ok := out.Cid["/"]; ok {
		msgCid = out.Cid["/"]
	}

	return msgCid, nil
}

func (api *MessageAPI) Status(ctx context.Context, msgCid string) (iface.MsgStatus, error) {
	var out iface.MsgStatus

	err := api.core().Request("message/status", msgCid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MessageAPI) Wait(ctx context.Context, msgCid string, opts ...fopt.MessageWaitOption) (iface.MsgWaitResult, error) {
	var out iface.MsgWaitResult

	options, err := fopt.MessageWaitOptions(opts...)
	if err != nil {
		return out, err
	}
	err = api.core().Request("message/wait", msgCid).
		Option("message", options.Message).
		Option("receipt", options.Receipt).
		Option("return", options.Return).
		Option("timeout", options.Timeout).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MessageAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

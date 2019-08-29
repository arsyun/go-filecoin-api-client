package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type ShowAPI HttpApi

func (api *ShowAPI) Block(ctx context.Context, headerCid string) (iface.ShowBlockInfo, error) {
	var out iface.ShowBlockInfo

	err := api.core().Request("show/block", headerCid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ShowAPI) Header(ctx context.Context, cid string) (iface.BlockMsg, error) {
	var out iface.BlockMsg

	err := api.core().Request("show/header", cid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ShowAPI) Messages(ctx context.Context, cid string) ([]iface.MessageDetail, error) {
	var out []iface.MessageDetail

	err := api.core().Request("show/messages", cid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ShowAPI) Receipts(ctx context.Context, cid string) ([]iface.ReceiptMsg, error) {
	var out []iface.ReceiptMsg

	err := api.core().Request("show/receipts", cid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ShowAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

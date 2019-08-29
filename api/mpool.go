package api

import (
	"context"
	iface "go-filecoin-api-client/interface-go-filecoin"
)

type MpoolAPI HttpApi

func (api *MpoolAPI) Ls(ctx context.Context) ([]iface.MessageDetail, error) {
	var out []iface.MessageDetail

	err := api.core().Request("mpool/ls").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MpoolAPI) Rm(ctx context.Context, msgid string) (error) {
	err := api.core().Request("mpool/rm").
		Arguments(msgid).
		Exec(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (api *MpoolAPI) Show(ctx context.Context, msgid string) (iface.MessageDetail, error) {
	var out iface.MessageDetail

	err := api.core().Request("mpool/show").
		Arguments(msgid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *MpoolAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

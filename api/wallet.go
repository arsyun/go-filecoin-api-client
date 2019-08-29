package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	iface "go-filecoin-api-client/interface-go-filecoin"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

type WalletAPI HttpApi

func (api *WalletAPI) Balance(ctx context.Context, address string) (string, error) {
	var out string

	err := api.core().Request("wallet/balance", address).
		Exec(ctx, &out)
	if err != nil {
		return "", err
	}

	return out, nil
}

func (api *WalletAPI) Export(ctx context.Context, addrs ...string) (iface.WalletInfo, error) {
	var out iface.WalletInfo

	//join multiple wallet address: arg=addr1&arg=addr2 
	values := make(url.Values)
	for _, arg := range addrs {
		values.Add("arg", arg)
	}

	err := api.core().Request(fmt.Sprintf("wallet/export?%s", values.Encode())).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *WalletAPI) Import(ctx context.Context, keyinfos []iface.KeyInfo) (addrs []string, err error) {
	var out struct {
		Addresses []string
	}
	w := make(map[string][]iface.KeyInfo)
	w["KeyInfo"] = keyinfos
	s, err := json.Marshal(w)
	if err != nil {
		return out.Addresses, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("walletFile", "")
	if _, err = io.Copy(formFile, strings.NewReader(string(s))); err != nil {
		return addrs, err
	}
	writer.Close()

	err = api.core().Request("wallet/import").
		Header("Content-Type", writer.FormDataContentType()).
		Body(body).
		Exec(ctx, &out)
	if err != nil {
		return out.Addresses, err
	}

	return out.Addresses, nil
}

func (api *WalletAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

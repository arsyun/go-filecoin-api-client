package api

import (
	"bytes"
	"context"
	"encoding/json"
	iface "go-filecoin-api-client/interface-go-filecoin"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

type ClientAPI HttpApi

func (api *ClientAPI) Cat(ctx context.Context, cid string) (io.Reader, error) {
	resp, err := api.core().Request("client/cat", cid).
		Send(ctx)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, err
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err := io.Copy(b, resp.Output); err != nil {
		return nil, err
	}

	return b, nil
}

func (api *ClientAPI) Import(ctx context.Context, fr io.Reader) (cid string, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", "")
	if _, err = io.Copy(formFile, fr); err != nil {
		return "", err
	}
	writer.Close()

	var out map[string]string
	err = api.core().Request("client/import").
		Header("Content-Type", writer.FormDataContentType()).
		Body(body).
		Exec(ctx, &out)
	if err != nil {
		return "", nil
	}
	if _, ok := out["/"]; ok {
		cid = out["/"]
	}

	return cid, nil
}

func (api *ClientAPI) ListAsks(ctx context.Context) ([]iface.Asks, error) {
	var out []iface.Asks

	resp, err := api.core().Request("client/list-asks").
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

	s := "[" + strings.TrimRight(strings.ReplaceAll(b.String(), "\n", ","), ",") + "]"
	if err := json.Unmarshal([]byte(s), &out); err != nil {
		return out, err
	}

	return out, nil
}

func (api *ClientAPI) Payments(ctx context.Context, dealCid string) ([]iface.PaymentsDeal, error) {
	var out []iface.PaymentsDeal

	err := api.core().Request("client/payments").
		Arguments(dealCid).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ClientAPI) ProposeStorageDeal(ctx context.Context, miner string, cid string, askId int, time int64) (iface.StorageDealInfo, error) {
	var out iface.StorageDealInfo

	err := api.core().Request("client/propose-storage-deal").
		Arguments(miner).
		Arguments(cid).
		Arguments(strconv.Itoa(askId)).
		Arguments(strconv.FormatInt(time, 10)).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *ClientAPI) QueryStorageDeal(ctx context.Context, dealID string) (iface.StorageDealInfo, error) {
	var out iface.StorageDealInfo

	err := api.core().Request("client/query-storage-deal", dealID).
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

//TODO 
func (api *ClientAPI) VerifyStorageDeal(ctx context.Context, dealID string) (iface.StorageDealVerify, error) {
	var out iface.StorageDealVerify
	return out, nil
}

func (api *ClientAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

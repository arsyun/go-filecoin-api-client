package api

import (
	"bytes"
	"context"
	"io"
)

type ChainAPI HttpApi

func (api *ChainAPI) Head(ctx context.Context) ([]string, error) {
	var out []map[string]string
	err := api.core().Request("chain/head").
		Exec(ctx, &out)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, v := range out {
		if _, ok := v["/"]; ok {
			res = append(res, v["/"])
		}
	}

	return res, nil
}

func (api *ChainAPI) Ls(ctx context.Context) (string, error) {
	resp, err := api.core().Request("chain/ls").
		Send(ctx)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	_, err = io.Copy(b, resp.Output)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (api *ChainAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

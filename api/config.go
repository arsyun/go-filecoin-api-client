package api

import (
	"bytes"
	"context"
	"io"
)

type ConfigAPI HttpApi

func (api *ConfigAPI) Get(ctx context.Context, key string) (string, error) {
	resp, err := api.core().Request("config").
		Arguments(key).
		Send(ctx)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err = io.Copy(b, resp.Output); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (api *ConfigAPI) Set(ctx context.Context, key, value string) (string, error) {
	resp, err := api.core().Request("config").
		Arguments(key).
		Arguments(value).
		Send(ctx)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err = io.Copy(b, resp.Output); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (api *ConfigAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

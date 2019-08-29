package api

import (
	"bytes"
	"context"
	fopt "go-filecoin-api-client/options"
	"io"
)

type LogAPI HttpApi

func (api *LogAPI) Level(ctx context.Context, level string, opts ...fopt.LogLevelOption) (string, error) {
	options, err := fopt.LogLevelOptions(opts...)
	if err != nil {
		return "", nil
	}

	req := api.core().Request("log/level", level)
	if options.Subsystem != "" {
		req = req.Option("subsystem", options.Subsystem)
	}

	resp, err := req.Send(ctx)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", resp.Error
	}

	b := new(bytes.Buffer)
	if _, err := io.Copy(b, resp.Output); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (api *LogAPI) Ls(ctx context.Context) ([]string, error) {
	var out []string
	err := api.core().Request("log/ls").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

//TODO
func (api *LogAPI) Tail(ctx context.Context) (string, error) {
	return "", nil
}

func (api *LogAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

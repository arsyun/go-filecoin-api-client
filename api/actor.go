package api

import (
	"bytes"
	"context"
	"encoding/json"
	iface "go-filecoin-api-client/interface-go-filecoin"
	"io"
	"strings"
)

type ActorAPI HttpApi

func (api *ActorAPI) Ls(ctx context.Context) ([]iface.AcotorInfo, error) {
	var out []iface.AcotorInfo
	resp, err := api.core().Request("actor/ls").
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

func (api *ActorAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

package api

import (
	"bytes"
	"context"
	"fmt"
	iface "go-filecoin-api-client/interface-go-filecoin"
	"io"
	"net/url"
	"strings"
)

type SwarmAPI HttpApi

func (api *SwarmAPI) Connect(ctx context.Context, addrs ...string) (connectIds []string, err error) {
	values := make(url.Values)
	for _, arg := range addrs {
		values.Add("arg", arg)
	}
	resp, err := api.core().Request(fmt.Sprintf("swarm/connect?%s", values.Encode())).
		Send(ctx)
	if err != nil {
		return connectIds, nil
	}
	if resp.Error != nil {
		return connectIds, resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err := io.Copy(b, resp.Output); err != nil {
		return connectIds, err
	}

	connectIds = strings.Split(strings.TrimRight(b.String(), "\n"), "\n")

	return connectIds, nil
}

func (api *SwarmAPI) Peers(ctx context.Context) (iface.SwarmPeers, error) {
	var out iface.SwarmPeers

	err := api.core().Request("swarm/peers").
		Exec(ctx, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (api *SwarmAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

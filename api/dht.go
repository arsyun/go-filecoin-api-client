package api

import (
	"bytes"
	"context"
	fopt "go-filecoin-api-client/options"
	"io"
	"strings"
)

type DhtAPI HttpApi

func (api *DhtAPI) FindPeer(ctx context.Context, peerID string) (addrs []string, err error) {
	resp, err := api.core().Request("dht/findpeer", peerID).
		Send(ctx)
	if err != nil {
		return addrs, err
	}
	if resp.Error != nil {
		return addrs, resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err = io.Copy(b, resp.Output); err != nil {
		return addrs, err
	}
	addrs = strings.Split(strings.TrimRight(b.String(), "\n"), "\n")

	return addrs, nil
}

func (api *DhtAPI) FindProvs(ctx context.Context, ID string, opts ...fopt.DhtFindProvsOption) (string, error) {
	options, err := fopt.DhtProvsOptions(opts...)
	if err != nil {
		return "", err
	}

	resp, err := api.core().Request("dht/findprovs").
		Option("num-providers", options.NumProviders).
		Arguments(ID).
		Send(ctx)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", resp.Error
	}
	defer resp.Close()

	b := new(bytes.Buffer)
	if _, err := io.Copy(b, resp.Output); err != nil {
		return "", err
	}

	return b.String(), nil
}

//TODO
func (api *DhtAPI) Query(ctx context.Context, peerID string) (string, error) {
	// err := api.core().Request("dht/query", peerID).
	// 	Exec(ctx, nil)
	// if err != nil {
	// 	return "", err
	// }
	return "", nil
}

func (api *DhtAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

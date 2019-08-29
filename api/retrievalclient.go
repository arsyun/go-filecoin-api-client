package api

import (
	"bytes"
	"context"
	"io"
)

type RetrievalClientAPI HttpApi

func (api *RetrievalClientAPI) RetrievePiece(ctx context.Context, minerAddr string, Cid string) (io.Reader, error) {
	resp, err := api.core().Request("retrieval-client/retrieve-piece", minerAddr, Cid).
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

func (api *RetrievalClientAPI) core() *HttpApi {
	return (*HttpApi)(api)
}

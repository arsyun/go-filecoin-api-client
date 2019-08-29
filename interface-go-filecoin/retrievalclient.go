package iface

import(
	"context"
	"io"
)

type RetrievalClientAPI interface {
	RetrievePiece(ctx context.Context, minerAddr, Cid string) (io.Reader, error)
}
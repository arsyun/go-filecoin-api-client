package iface

import (
	"context"
	"go-filecoin-api-client/options"
)

type DhtAPI interface {
	FindPeer(ctx context.Context, peerID string) (multiaddresses []string, err error)

	FindProvs(ctx context.Context, ID string, opts ...options.DhtFindProvsOption) (string, error)
	//TODO
	Query(ctx context.Context, peerID string) (string, error)
}

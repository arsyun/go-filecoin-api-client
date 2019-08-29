package iface

import(
	"context"
)

type ShowBlockInfo struct {
	Header		BlockMsg		`json:"header"`
	Messages	[]MessageDetail	`json:"messages"`
	Receipts    []ReceiptMsg	`json:"receipts"`
}


type ShowAPI interface {
	Block(ctx context.Context, headerCid string) (ShowBlockInfo, error)

	Header(ctx context.Context, cid string) (BlockMsg, error)

	Messages(ctx context.Context, cid string) ([]MessageDetail, error)

	Receipts(ctx context.Context, cid string) ([]ReceiptMsg, error)
}
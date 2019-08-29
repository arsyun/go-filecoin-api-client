package iface

import (
	"context"
	fopt "go-filecoin-api-client/options"
)

type Paych struct {
	Target         string      `json:"target"`
	Amount         string      `json:"amount"`
	AmountRedeemed string      `json:"amount_redeemed"`
	AgreedEol      int64       `json:"agreed_eol"`
	Condition      interface{} `json:"condition"`
	Eol            int64       `json:"eol"`
	Redeemed       bool        `json:"redeemed"`
}

type ChannelInfo struct {
	Cid     map[string]string `json:"cid"`
	GasUsed string            `json:"gas_used"`
	Preview bool              `json:"preview"`
}

type PaychAPI interface {
	Cancel(ctx context.Context, channelID int64, opts ...fopt.PaychOption) (ChannelInfo, error)

	Close(ctx context.Context, voucher string, opts ...fopt.PaychOption) (ChannelInfo, error)

	Create(ctx context.Context, target string, amount string, eol int64, opts ...fopt.PaychOption) (ChannelInfo, error)

	Extend(ctx context.Context, channelID int64, amount string, eol int64, opts ...fopt.PaychOption) (ChannelInfo, error)

	Ls(context.Context) (map[string]Paych, error)

	Reclaim(ctx context.Context, channelID int64, opts ...fopt.PaychOption) (ChannelInfo, error)

	Redeem(ctx context.Context, voucher string, opts ...fopt.PaychOption) (ChannelInfo, error)

	Voucher(ctx context.Context, channelID int64, amount string, opts ...fopt.PaychVoucherOption) (voucher string, err error)
}

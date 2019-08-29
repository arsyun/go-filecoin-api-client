package iface

import (
	"context"
	fopt "go-filecoin-api-client/options"
)

type MsgSendResult struct {
	Cid     map[string]string `json:"cid"`
	GasUsed string            `json:"gas_used"`
	Preview bool              `json:"preview"`
}

type MsgStatus struct {
	InPool    bool     `json:"in_pool"`
	PoolMsg   string   `json:"pool_msg"`
	InOutbox  bool     `json:"in_outbox"`
	OutboxMsg string   `json:"outbox_msg"`
	ChainMsg  ChainMsg `json:"chain_msg"`
}

type MsgWaitResult struct {
	Message   MessageDetail `json:"message"`
	Receipt   ReceiptMsg    `json:"receipt"`
	Signature string        `json:"signature"`
}

type ChainMsg struct {
	Message MessageDetail `json:"message"`
	Block   BlockMsg      `json:"block"`
	Receipt ReceiptMsg    `json:"receipt"`
}

type MessageDetail struct {
	MeteredMessage MeteredMessage `json:"metered_message"`
	Signature      string         `json:"signature"`
}

type MeteredMessage struct {
	Message  MessageInfo `json:"message"`
	GasPrice string      `json:"gas_price"`
	GasLimit string      `json:"gas_limit"`
}

type MessageInfo struct {
	To     string `json:"to"`
	From   string `json:"from"`
	Nonce  string `json:"nonce"`
	Value  string `json:"value"`
	Method string `json:"method"`
	Params string `json:"params"`
}

type BlockMsg struct {
	Miner           string            `json:"miner"`
	Ticket          string            `json:"ticket"`
	Parents         map[string]string `json:"parents"`
	ParentWeight    string            `json:"parent_weight"`
	Height          string            `json:"height"`
	Nonce           string            `json:"nonce"`
	Messages        map[string]string `json:"messages"`
	StateRoot       map[string]string `json:"state_root"`
	MessageReceipts map[string]string `json:"message_receipts"`
	Proof           string            `json:"proof"`
	TimeStamp       string            `json:"time_stamp"`
}

type ReceiptMsg struct {
	ExitCode   int      `json:"exit_code"`
	Return     []string `json:"return"`
	GasAttoFIL string   `json:"gas_atto_fil"`
}

type MessageAPI interface {
	Send(ctx context.Context, target string, method string, opts ...fopt.MessageSendOption) (msgCid string, err error)

	Status(ctx context.Context, msgCid string) (MsgStatus, error)

	Wait(ctx context.Context, msgCid string, opts ...fopt.MessageWaitOption) (MsgWaitResult, error)
}

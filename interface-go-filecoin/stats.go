package iface

import (
	"context"
)

type StatsBandWidth struct {
	TotalIn  int64   `json:"total_in"`
	TotalOut int64   `json:"total_out"`
	RateIn   float64 `json:"rate_in"`
	RateOut  float64 `json:"rate_out"`
}

type StatsAPI interface {
	BandWidth(ctx context.Context) (StatsBandWidth, error)
}

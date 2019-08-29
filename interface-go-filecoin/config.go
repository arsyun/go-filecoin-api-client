package iface

import (
	"context"
)

type Config struct {
	Api           ApiConfig           `json:"api"`
	Bootstrap     BootstrapConfig     `json:"bootstrap"`
	Datastore     DatastoreConfig     `json:"datastore"`
	Heartbeat     HeartbeatConfig     `json:"heartbeat"`
	Mpool         MpoolConfig         `json:"mpool"`
	Mining        MiningCofig         `json:"mining"`
	Net           string              `json:"net"`
	Observability ObservabilityConfig `json:"observability"`
	Sectorbase    SectorbaseConfig    `json:"sectorbase"`
	Swarm         SwarmConfig         `json:"swarm"`
	Wallet        WalletConfig        `json:"wallet"`
}

type ApiConfig struct {
	Address                       string   `json:"address"`
	AccessControlAllowOrigin      []string `json:"accessControlAlloworigin"`
	AccessControlAllowCredentials bool     `json:"accessControlAllowcredentials"`
	AccessControlAllowMethods     []string `json:"accessControlAllowmethods"`
}

type DatastoreConfig struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type HeartbeatConfig struct {
	BeatTarget      string `json:"beat_target"`
	BeatPeriod      string `json:"beat_period"`
	ReconnectPeriod string `json:"reconnect_period"`
	Nickname        string `json:"nickname"`
}

type SectorbaseConfig struct {
	Rootdir string `json:"rootdir"`
}

type ObservabilityConfig struct {
	Metrics MetricsConfig `json:"metrics"`
	Tracing TracingConfig `json:"tracing"`
}

type MetricsConfig struct {
	PrometheusEnabled  bool   `json:"prometheus_enabled"`
	ReportInterval     string `json:"report_interval"`
	PrometheusEndpoint string `json:"prometheus_endpoint"`
}

type TracingConfig struct {
	JaegerTracingEnabled bool   `json:"jaeger_tracing_enabled"`
	ProbabilitySampler   int32  `json:"probability_sampler"`
	JaegerEndpoint       string `json:"jaeger_endpoint"`
}

type ConfigAPI interface {
	Get(ctx context.Context, key string) (string, error)

	Set(ctx context.Context, key, value string) (string, error)
}

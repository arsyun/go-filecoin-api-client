package iface

import (
	"context"
)

type InspectAll struct {
	Config          Config      `json:"config"`
	Runtime         Runtime     `json:"runtime"`
	Environment     Environment `json:"environment"`
	Disk            Disk        `json:"disk"`
	Memory          Memory      `json:"memory"`
	FilecoinVersion string      `json:"filecoin_version"`
}

type Runtime struct {
	Os            string `json:"os"`
	Arch          string `json:"arch"`
	Version       string `json:"version"`
	Compiler      string `json:"compiler"`
	NumProc       int32  `json:"num_proc"`
	GoMaxProcs    int32  `json:"go_max_procs"`
	NumGoRoutines int64  `json:"num_go_routines"`
	NumCGoCalls   int64  `json:"num_cgo_calls"`
}

type Environment struct {
	FILAPI  string `json:"filapi"`
	FILPATH string `json:"filpath"`
	GOPATH  string `json:"gopath"`
}

type Disk struct {
	Free   int64  `json:"free"`
	Total  int64  `json:"total"`
	FSType string `json:"fstype"`
}

type Memory struct {
	Swap    int   `json:"swap"`
	Virtual int64 `json:"virtual"`
}

type InspectAPI interface {
	All(context.Context) (InspectAll, error)

	Config(context.Context) (Config, error)

	Disk(context.Context) (Disk, error)

	Environment(context.Context) (Environment, error)

	Memory(context.Context) (Memory, error)

	Runtime(context.Context) (Runtime, error)
}

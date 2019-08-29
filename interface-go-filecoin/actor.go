package iface

import (
	"context"
)

type AcotorInfo struct {
	Actortype string            `json:"actorType"`
	Address   string            `json:"address"`
	Code      map[string]string `json:"code"`
	nonce     int64             `json:"nonce"`
	Balance   string            `json:"balance"`
	Exports   Exports           `json:"exports"`
	Head      map[string]string `json:"head"`
}

type Exports struct {
	AddAsk                   BasicExportsData `json:"add_ask"`
	CalculateLateFee         BasicExportsData `json:"calculate_late_fee"`
	ChangeWorker             BasicExportsData `json:"change_worker"`
	CommitSector             BasicExportsData `json:"commit_sector"`
	GetActiveCollateral      BasicExportsData `json:"get_active_collateral"`
	GetAsk                   BasicExportsData `json:"get_ask"`
	GetAsks                  BasicExportsData `json:"get_asks"`
	GetLastUsedSectorID      BasicExportsData `json:"get_last_used_sector_id"`
	GetOwner                 BasicExportsData `json:"getOwner"`
	GetPeerID                BasicExportsData `json:"get_peerID"`
	GetPoStState             BasicExportsData `json:"get_post_state"`
	GetPower                 BasicExportsData `json:"get_power"`
	GetProvingPeriod         BasicExportsData `json:"get_proving_period"`
	GetProvingSetCommitments BasicExportsData `json:"get_proving_set_commitments"`
	GetSectorSize            BasicExportsData `json:"get_sector_size"`
	GetWorker                BasicExportsData `json:"get_worker"`
	IsBootstrapMiner         BasicExportsData `json:"is_bootstrap_miner"`
	SlashStorageFault        BasicExportsData `json:"slash_storage_fault"`
	SubmitPoSt               BasicExportsData `json:"submit_po_st"`
	UpdatePeerID             BasicExportsData `json:"update_peerID"`
	VerifyPieceInclusion     BasicExportsData `json:"verify_piece_inclusion"`
}

type BasicExportsData struct {
	Params []string
	Return []string
}

type ActorAPI interface {
	Ls(context.Context) ([]AcotorInfo, error)
}

package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
)

// QcrescentQuery Contains Qcrescent custom queries
type QcrescentQuery struct {
	// Query state of Qcrescent keeper to find the balance of account of other chain
	QueryState *QueryQueryStateRequest `json:"query_query_state_request"`
}

type QueryQueryStateRequest struct {
	Sequence uint64 `json:"sequence"`
}
type QueryQueryStateResponse struct {
	// Address wasmvmtypes.HumanAddress        `json:"address"`
	Coins wasmvmtypes.Coins `json:"coins"`
}

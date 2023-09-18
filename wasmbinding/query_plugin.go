package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	wasmbinding "github.com/placeholder-dapps/athena/wasmbinding/bindings"
)

func CustomQuerier(qk *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery wasmbinding.QcrescentQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "Qcresent Query")
		}

		if contractQuery.QueryState != nil {
			stateResp, err := qk.GetQueryState(ctx, contractQuery.QueryState.Sequence)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "query state for balance")
			}

			bz, err := json.Marshal(stateResp)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "while marshaling state response")
			}
			return bz, nil
		}
		return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknow qcrescent query"}
	}
}

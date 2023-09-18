package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	wasmbinding "github.com/placeholder-dapps/athena/wasmbinding/bindings"
	qcrescentKeeper "github.com/placeholder-dapps/athena/x/qcrescent/keeper"
)

type QueryPlugin struct {
	qcrescentKeeper *qcrescentKeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin
func NewQueryPlugin(qk *qcrescentKeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		qcrescentKeeper: qk,
	}
}

func (qk QueryPlugin) GetQueryState(ctx sdk.Context, sequence uint64) (*wasmbinding.QueryQueryStateResponse, error) {
	queryResponse, err := qk.qcrescentKeeper.GetQueryResponse(ctx, sequence)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "getting query response")

	}
	var response wasmbinding.QueryQueryStateResponse
	response = wasmbinding.QueryQueryStateResponse{
		Coins: ConvertSdkCoinsToWasmCoins(queryResponse.Balances),
	}
	return &response, nil
}

// ConvertSdkCoinsToWasmCoins converts sdk type coins to wasm vm type coins
func ConvertSdkCoinsToWasmCoins(coins sdk.Coins) wasmvmtypes.Coins {
	var toSend wasmvmtypes.Coins
	for _, coin := range coins {
		c := ConvertSdkCoinToWasmCoin(coin)
		toSend = append(toSend, c)
	}
	return toSend
}

// ConvertSdkCoinToWasmCoin converts a sdk type coin to a wasm vm type coin
func ConvertSdkCoinToWasmCoin(coin sdk.Coin) wasmvmtypes.Coin {
	return wasmvmtypes.Coin{
		Denom: coin.Denom,
		// Note: gamm tokens have 18 decimal places, so 10^22 is common, no longer in u64 range
		Amount: coin.Amount.String(),
	}
}

package wasmbinding

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	wasmbinding "github.com/placeholder-dapps/athena/wasmbinding/bindings"
	qcrescentKeeper "github.com/placeholder-dapps/athena/x/qcrescent/keeper"
	qcrescenttypes "github.com/placeholder-dapps/athena/x/qcrescent/types"
)

type CustomMessenger struct {
	wrapped         wasmkeeper.Messenger
	qcrescentKeeper *qcrescentKeeper.Keeper
}

// DispatchMsg implements keeper.Messenger.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) (events []sdk.Event, data [][]byte, err error) {
	if msg.Custom != nil {
		var contractMsg wasmbinding.QcrescentMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "qcrescent msg")
		}
		if contractMsg.SendQueryAllBalance != nil {
			return m.sendQueryAllBalance(ctx, contractAddr, contractMsg.SendQueryAllBalance)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func CustomMessageDecorator(qcrescent *qcrescentKeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:         old,
			qcrescentKeeper: qcrescent,
		}
	}
}
func (m *CustomMessenger) sendQueryAllBalance(ctx sdk.Context, contractAddr sdk.AccAddress, queryAllBalances *wasmbinding.SendQueryAllBalance) ([]sdk.Event, [][]byte, error) {
	err := PerformSendQueryAllBalances(m.qcrescentKeeper, ctx, contractAddr, queryAllBalances)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform mint")
	}
	return nil, nil, nil
}
func PerformSendQueryAllBalances(f *qcrescentKeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, sendQuery *wasmbinding.SendQueryAllBalance) error {
	if sendQuery != nil {
		return wasmvmtypes.InvalidRequest{Err: "send query all balance is null"}
	}
	queryAddress, err := parseAddress(sendQuery.Address)
	if err != nil {
		return err
	}

	msgServer := qcrescentKeeper.NewMsgServerImpl(*f)
	_, err = msgServer.MsgSendQueryAllBalances(ctx, qcrescenttypes.NewMsgSendQueryAllBalances(sendQuery.Creator, sendQuery.ChannelId, queryAddress.String(), sendQuery.Pagination))
	if err != nil {
		return sdkerrors.Wrap(err, "querying balance of account on other chain")
	}
	return nil
}
func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "verify address format")
	}
	return parsed, nil
}

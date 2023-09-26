package simulation

import (
	"math/rand"

	"github.com/aerius-labs/athena/x/qcrescent/keeper"
	"github.com/aerius-labs/athena/x/qcrescent/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMsgSendQueryAllBalances(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMsgSendQueryAllBalances{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MsgSendQueryAllBalances simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MsgSendQueryAllBalances simulation not implemented"), nil, nil
	}
}

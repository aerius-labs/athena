package qcrescent

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/placeholder-dapps/athena/testutil/sample"
	qcrescentsimulation "github.com/placeholder-dapps/athena/x/qcrescent/simulation"
	"github.com/placeholder-dapps/athena/x/qcrescent/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = qcrescentsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgMsgSendQueryAllBalances = "op_weight_msg_msg_send_query_all_balances"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMsgSendQueryAllBalances int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	qcrescentGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&qcrescentGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMsgSendQueryAllBalances int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMsgSendQueryAllBalances, &weightMsgMsgSendQueryAllBalances, nil,
		func(_ *rand.Rand) {
			weightMsgMsgSendQueryAllBalances = defaultWeightMsgMsgSendQueryAllBalances
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMsgSendQueryAllBalances,
		qcrescentsimulation.SimulateMsgMsgSendQueryAllBalances(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgMsgSendQueryAllBalances,
			defaultWeightMsgMsgSendQueryAllBalances,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				qcrescentsimulation.SimulateMsgMsgSendQueryAllBalances(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

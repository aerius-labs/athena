package qcrescent_test

import (
	"testing"

	keepertest "github.com/aerius-labs/athena/testutil/keeper"
	"github.com/aerius-labs/athena/testutil/nullify"
	"github.com/aerius-labs/athena/x/qcrescent"
	"github.com/aerius-labs/athena/x/qcrescent/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QcrescentKeeper(t)
	qcrescent.InitGenesis(ctx, *k, genesisState)
	got := qcrescent.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}

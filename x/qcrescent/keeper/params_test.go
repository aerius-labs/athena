package keeper_test

import (
	"testing"

	testkeeper "github.com/aerius-labs/athena/testutil/keeper"
	"github.com/aerius-labs/athena/x/qcrescent/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QcrescentKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

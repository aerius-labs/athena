package keeper_test

import (
	"testing"

	testkeeper "github.com/placeholder-dapps/athena/testutil/keeper"
	"github.com/placeholder-dapps/athena/x/qcrescent/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QcrescentKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/aerius-labs/athena/testutil/keeper"
	"github.com/aerius-labs/athena/x/qcrescent/keeper"
	"github.com/aerius-labs/athena/x/qcrescent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.QcrescentKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

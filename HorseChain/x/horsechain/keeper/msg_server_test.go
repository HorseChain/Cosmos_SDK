package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/HorseChain/HorseChain/testutil/keeper"
	"github.com/HorseChain/HorseChain/x/horsechain/keeper"
	"github.com/HorseChain/HorseChain/x/horsechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HorsechainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

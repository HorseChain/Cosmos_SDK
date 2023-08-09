package keeper_test

import (
	"testing"

	testkeeper "github.com/HorseChain/HorseChain/testutil/keeper"
	"github.com/HorseChain/HorseChain/x/horsechain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HorsechainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

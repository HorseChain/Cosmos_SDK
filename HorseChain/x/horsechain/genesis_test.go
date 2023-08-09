package horsechain_test

import (
	"testing"

	keepertest "github.com/HorseChain/HorseChain/testutil/keeper"
	"github.com/HorseChain/HorseChain/testutil/nullify"
	"github.com/HorseChain/HorseChain/x/horsechain"
	"github.com/HorseChain/HorseChain/x/horsechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HorsechainKeeper(t)
	horsechain.InitGenesis(ctx, *k, genesisState)
	got := horsechain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

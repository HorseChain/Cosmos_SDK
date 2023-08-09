package keeper

import (
	"github.com/HorseChain/HorseChain/x/horsechain/types"
)

var _ types.QueryServer = Keeper{}

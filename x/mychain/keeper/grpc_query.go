package keeper

import (
	"github.com/daniel/mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}

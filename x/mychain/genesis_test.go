package mychain_test

import (
	"testing"

	keepertest "github.com/daniel/mychain/testutil/keeper"
	"github.com/daniel/mychain/x/mychain"
	"github.com/daniel/mychain/x/mychain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MychainKeeper(t)
	mychain.InitGenesis(ctx, *k, genesisState)
	got := mychain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}

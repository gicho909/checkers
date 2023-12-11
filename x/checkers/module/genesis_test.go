package checkers_test

import (
	"testing"

	keepertest "github.com/gicho909/checkers/testutil/keeper"
	"github.com/gicho909/checkers/testutil/nullify"
	"github.com/gicho909/checkers/x/checkers/module"
	"github.com/gicho909/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, k, genesisState)
	got := checkers.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
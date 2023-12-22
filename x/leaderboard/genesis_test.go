package leaderboard_test

import (
	"testing"

	keepertest "github.com/gicho909/checkers/testutil/keeper"
	"github.com/gicho909/checkers/testutil/nullify"
	"github.com/gicho909/checkers/x/leaderboard"
	"github.com/gicho909/checkers/x/leaderboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LeaderboardKeeper(t)
	leaderboard.InitGenesis(ctx, *k, genesisState)
	got := leaderboard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

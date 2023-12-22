package keeper

import (
	"github.com/gicho909/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}

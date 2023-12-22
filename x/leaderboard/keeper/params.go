package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gicho909/checkers/x/leaderboard/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Length(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Length returns the Length param
func (k Keeper) Length(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyLength, &res)
	return
}

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gicho909/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestCreate3GamesHasSavedFifo(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	// Second game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	systemInfo2, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo2)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game2)

	// Third game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	systemInfo3, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo3)
	game1, found = keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found = keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "3",
	}, game2)
	game3, found := keeper.GetStoredGame(ctx, "3")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       alice,
		Red:         bob,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "2",
		AfterIndex:  "-1",
	}, game3)
}
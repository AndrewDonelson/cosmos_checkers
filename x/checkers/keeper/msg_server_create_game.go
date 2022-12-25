package keeper

import (
	"context"
	"strconv"

	"checkers/x/checkers/rules"
	"checkers/x/checkers/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the current systemInfo object using the Keeper.GetSystemInfo method
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	// Using the rules package, create a new game and store it in a StoredGame object
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: newIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}

	// Confirm that the values in the object are correct by checking the validity of the players' addresses
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	// Save the StoredGame object using the Keeper.SetStoredGame method
	k.Keeper.SetStoredGame(ctx, storedGame)

	// Prepare the ground for the next game using the Keeper.SetSystemInfo method to increment the NextId field and
	// save the updated systemInfo object
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	_ = ctx

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}

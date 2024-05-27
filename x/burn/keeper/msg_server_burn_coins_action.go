package keeper

import (
	"context"

	"github.com/BitCannaGlobal/bcna/x/burn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BurnCoinsAction(goCtx context.Context, msg *types.MsgBurnCoinsAction) (*types.MsgBurnCoinsActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnCoinsActionResponse{}, nil
}

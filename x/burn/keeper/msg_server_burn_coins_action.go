package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/BitCannaGlobal/bcna/x/burn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Move coins from sender to Bank account module and then the module burns the coins.
func (k msgServer) BurnCoinsAction(goCtx context.Context, msg *types.MsgBurnCoinsAction) (*types.MsgBurnCoinsActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address: %s", err)
	}

	// Validate the coins
	coins := sdk.NewCoins(msg.Amount)
	if !coins.IsValid() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// Get the module's params to verify the allowed denom
	params := k.GetParams(ctx)
	if msg.Amount.Denom != params.BurnDenom {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("denomination mismatch: expected %s, got %s", params.BurnDenom, msg.Amount.Denom)
	}
	// Ensure coins are positive
	if !coins.IsAllPositive() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// Gets the balance of the sender to check if are there enough coins.
	balance := k.bankKeeper.GetBalance(ctx, creatorAddr, msg.Amount.Denom)
	if balance.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("insufficient balance for %s", creatorAddr)
	}

	// Send the coins from the creator to the module and later it burns
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("failed to send coins from account to module: %v", err)
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("failed to burn coins: %v", err)
	}
	// Log the successful burn operation
	k.Logger(ctx).Info("Burning coins!! ", "signer", msg.Creator, "amount", msg.Amount)

	return &types.MsgBurnCoinsActionResponse{}, nil
}

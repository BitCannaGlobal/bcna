package keeper

import (
	"context"
	"fmt"

	"github.com/BitCannaGlobal/bcna/x/burn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Move coins from sender to Bank account module and then the module burns the coins.
func (k msgServer) BurnCoinsAction(goCtx context.Context, msg *types.MsgBurnCoinsAction) (*types.MsgBurnCoinsActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("the address is not valid")
	}

	// Check nulls and valid amounts
	if msg.Coins == nil || len(msg.Coins) == 0 {
		return nil, fmt.Errorf("no coins specified or the amount is not valid")
	}
	// for _, coin := range msg.Coins {
	// 	if coin.Amount.LTE(sdk.ZeroInt()) { // Comprueba si la cantidad es menor o igual a cero.
	// 		return nil, fmt.Errorf("invalid amount for coin %s, amount must be positive", coin.Denom)
	// 	}
	// }

	// Gets the balance of the sender to check if are there enough coins.
	for _, coin := range msg.Coins {
		balance := k.bankKeeper.GetBalance(ctx, creatorAddr, coin.Denom)
		if balance.Amount.LT(coin.Amount) {
			return nil, fmt.Errorf("insufficient balance for coin %s", coin.Denom)
		}
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, msg.Coins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Coins); err != nil {
		return nil, err
	}
	return &types.MsgBurnCoinsActionResponse{}, nil
}

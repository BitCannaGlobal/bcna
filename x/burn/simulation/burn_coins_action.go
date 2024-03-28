package simulation

import (
	"math/rand"

	"github.com/BitCannaGlobal/bcna/x/burn/keeper"
	"github.com/BitCannaGlobal/bcna/x/burn/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// func SimulateMsgBurnCoinsAction(
// 	ak types.AccountKeeper,
// 	bk types.BankKeeper,
// 	k keeper.Keeper,
// ) simtypes.Operation {
// 	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)
// 		msg := &types.MsgBurnCoinsAction{
// 			Creator: simAccount.Address.String(),
// 		}

// 		// TODO: Handling the BurnCoinsAction simulation

//			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BurnCoinsAction simulation not implemented"), nil, nil
//		}
//	}
func SimulateMsgBurnCoinsAction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	cdc *codec.ProtoCodec,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// Get the balance of the account
		balance := bk.GetBalance(ctx, simAccount.Address, "ubcna")

		// Has the account coins?
		if balance.IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgBurnCoinsAction, "Balance insuficiente"), nil, nil
		}

		// Assign a ramdom amount (within the max balance)
		amount := simtypes.RandIntBetween(r, 1, int(balance.Amount.Int64()))

		// Build and send the message
		msg := &types.MsgBurnCoinsAction{
			Creator: simAccount.Address.String(),
			Amount:  sdk.NewCoin(balance.Denom, sdk.NewInt(int64(amount))),
		}

		opMsg := simtypes.NewOperationMsg(msg, true, "", cdc)

		return opMsg, nil, nil
	}
}

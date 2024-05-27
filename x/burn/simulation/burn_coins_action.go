package simulation

import (
	"math/rand"

	"github.com/BitCannaGlobal/bcna/x/burn/keeper"
	"github.com/BitCannaGlobal/bcna/x/burn/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBurnCoinsAction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnCoinsAction{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BurnCoinsAction simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "BurnCoinsAction simulation not implemented"), nil, nil
	}
}

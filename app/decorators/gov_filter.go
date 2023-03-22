package decorators

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var MiniumInitialDepositRate = sdk.NewDecWithPrec(20, 2)

type GovPreventSpamDecorator struct {
	govKeeper *govkeeper.Keeper
	cdc       codec.BinaryCodec
}

func NewGovPreventSpamDecorator(cdc codec.BinaryCodec, govKeeper *govkeeper.Keeper) GovPreventSpamDecorator {
	return GovPreventSpamDecorator{
		govKeeper: govKeeper,
		cdc:       cdc,
	}
}

func (gpsd GovPreventSpamDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx,
	simulate bool, next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	// run checks only on CheckTx or simulate
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}
	msgs := tx.GetMsgs()

	err = gpsd.checkSpamSubmitProposalMsg(ctx, msgs)

	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

// validateGovMsgs checks if the InitialDeposit amounts are greater than the minimum initial deposit amount
func (gpsd GovPreventSpamDecorator) checkSpamSubmitProposalMsg(ctx sdk.Context, msgs []sdk.Msg) error {
	// prevent spam gov msg
	depositParams := gpsd.govKeeper.GetDepositParams(ctx)
	miniumInitialDeposit := gpsd.calcMiniumInitialDeposit(depositParams.MinDeposit)

	validMsg := func(m sdk.Msg) error {
		switch msg := m.(type) {
		case *govv1beta1.MsgSubmitProposal:
			// // prevent spam gov msg

			if msg.InitialDeposit.IsAllLT(miniumInitialDeposit) {
				return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "not enough initial deposit. required: %v", miniumInitialDeposit)
			}
		case *govv1.MsgSubmitProposal:
			// // prevent spam gov msg at v1

			// if msg.InitialDeposit.IsAllLT(miniumInitialDeposit) {
			// 	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "not enough initial deposit. required: %v", miniumInitialDeposit)
			// }
			// panic: don't use Gov v1 messages:
			message := "- Please don't use Gov v1 in SDK v0.46! "
			panic(fmt.Errorf("failed to create AnteHandler: %s", message))

		}

		return nil
	}

	validAuthz := func(execMsg *authz.MsgExec) error {
		// depositParams := gpsd.govKeeper.GetDepositParams(ctx)
		// miniumInitialDeposit := gpsd.calcMiniumInitialDeposit(depositParams.MinDeposit)
		for _, v := range execMsg.Msgs {
			var innerMsg sdk.Msg
			err := gpsd.cdc.UnpackAny(v, &innerMsg)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "not enough initial deposit. required: %v", miniumInitialDeposit)
			}

			err = validMsg(innerMsg)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, m := range msgs {
		if msg, ok := m.(*authz.MsgExec); ok {
			if err := validAuthz(msg); err != nil {
				return err
			}
			continue
		}

		// validate normal msgs
		err := validMsg(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gpsd GovPreventSpamDecorator) calcMiniumInitialDeposit(minDeposit sdk.Coins) (miniumInitialDeposit sdk.Coins) {
	for _, coin := range minDeposit {
		miniumInitialCoin := MiniumInitialDepositRate.MulInt(coin.Amount).RoundInt()
		miniumInitialDeposit = miniumInitialDeposit.Add(sdk.NewCoin(coin.Denom, miniumInitialCoin))
	}

	return
}

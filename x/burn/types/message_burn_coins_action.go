package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnCoinsAction{}

func NewMsgBurnCoinsAction(creator string, coins sdk.Coins) *MsgBurnCoinsAction {
	return &MsgBurnCoinsAction{
		Creator: creator,
		Coins:   coins,
	}
}

func (msg *MsgBurnCoinsAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

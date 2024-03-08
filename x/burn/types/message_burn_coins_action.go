package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnCoinsAction = "burn_coins_action"

var _ sdk.Msg = &MsgBurnCoinsAction{}

func NewMsgBurnCoinsAction(creator string, coins sdk.Coins) *MsgBurnCoinsAction {
	return &MsgBurnCoinsAction{
		Creator: creator,
		Coins:   coins,
	}
}

func (msg *MsgBurnCoinsAction) Route() string {
	return RouterKey
}

func (msg *MsgBurnCoinsAction) Type() string {
	return TypeMsgBurnCoinsAction
}

func (msg *MsgBurnCoinsAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnCoinsAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnCoinsAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

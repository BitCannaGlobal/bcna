package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
		return fmt.Errorf("invalid creator address: %v: %w", err, errors.New("invalid address"))
	}
	return nil
}

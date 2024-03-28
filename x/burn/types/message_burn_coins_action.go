package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgBurnCoinsAction = "burn_coins_action"

var _ sdk.Msg = &MsgBurnCoinsAction{}

func NewMsgBurnCoinsAction(creator string, amount sdk.Coin) *MsgBurnCoinsAction {
	return &MsgBurnCoinsAction{
		Creator: creator,
		Amount:  amount,
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
	if msg.Amount.IsNegative() || msg.Amount.IsZero() || !msg.Amount.IsValid() {
		return fmt.Errorf("amount must be positive or valid")
	}
	// Let's filter the length in the amount (max digits or bigger possible number )
	maxDigits := 18
	if len(msg.Amount.Amount.String()) > maxDigits {
		return fmt.Errorf("token amount exceeds maximum length of %d digits", maxDigits)
	}
	return nil
}

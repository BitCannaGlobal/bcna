package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBitcannaid{}

func NewMsgCreateBitcannaid(creator string, bcnaid string, address string) *MsgCreateBitcannaid {
	return &MsgCreateBitcannaid{
		Creator: creator,
		Bcnaid:  bcnaid,
		Address: address,
	}
}

func (msg *MsgCreateBitcannaid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBitcannaid{}

func NewMsgUpdateBitcannaid(creator string, id uint64, bcnaid string, address string) *MsgUpdateBitcannaid {
	return &MsgUpdateBitcannaid{
		Id:      id,
		Creator: creator,
		Bcnaid:  bcnaid,
		Address: address,
	}
}

func (msg *MsgUpdateBitcannaid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBitcannaid{}

func NewMsgDeleteBitcannaid(creator string, id uint64) *MsgDeleteBitcannaid {
	return &MsgDeleteBitcannaid{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteBitcannaid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

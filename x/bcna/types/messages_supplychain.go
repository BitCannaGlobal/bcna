package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSupplychain{}

func NewMsgCreateSupplychain(creator string, product string, info string, supplyinfo string, supplyextra string) *MsgCreateSupplychain {
	return &MsgCreateSupplychain{
		Creator:     creator,
		Product:     product,
		Info:        info,
		Supplyinfo:  supplyinfo,
		Supplyextra: supplyextra,
	}
}

func (msg *MsgCreateSupplychain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSupplychain{}

func NewMsgUpdateSupplychain(creator string, id uint64, product string, info string, supplyinfo string, supplyextra string) *MsgUpdateSupplychain {
	return &MsgUpdateSupplychain{
		Id:          id,
		Creator:     creator,
		Product:     product,
		Info:        info,
		Supplyinfo:  supplyinfo,
		Supplyextra: supplyextra,
	}
}

func (msg *MsgUpdateSupplychain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSupplychain{}

func NewMsgDeleteSupplychain(creator string, id uint64) *MsgDeleteSupplychain {
	return &MsgDeleteSupplychain{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteSupplychain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

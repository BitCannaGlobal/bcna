package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/bcna module sentinel errors
var (
	ErrInvalidSigner       = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrDuplicateBitcannaid = sdkerrors.Register(ModuleName, 1101, "BitCannaID already exists")
	ErrKeyNotFound         = sdkerrors.Register(ModuleName, 1102, "Key doesn't exists")
	ErrUnauthorized        = sdkerrors.Register(ModuleName, 1103, "Incorrect owner")
	ErrUnrecognized        = sdkerrors.Register(ModuleName, 1104, "Unrecognized messager")
	ErrInvalidAddress      = sdkerrors.Register(ModuleName, 1105, "invalid address")
	ErrMaxCharacters       = sdkerrors.Register(ModuleName, 1106, "input exceeds the permitted length limit")
)

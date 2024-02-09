package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/bcna module sentinel errors
var (
	ErrInvalidSigner       = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrDuplicateBitcannaid = sdkerrors.Register(ModuleName, 1101, "BitCannaID already exists")
	ErrMaxCharacters       = sdkerrors.Register(ModuleName, 1106, "input exceeds the permitted length limit")
)

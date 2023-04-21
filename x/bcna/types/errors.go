package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bcna module sentinel errors
var (
	ErrDuplicateBitcannaid = sdkerrors.Register(ModuleName, 1101, "BitCannaID already exists")
)

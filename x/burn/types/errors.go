package types

// DONTCOVER

import (
	sdkioerrors "cosmossdk.io/errors"
)

// x/burn module sentinel errors
var (
	ErrSample = sdkioerrors.Register(ModuleName, 1101, "Custom error")
)

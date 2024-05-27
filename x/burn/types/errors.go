package types

// DONTCOVER

import (
	sdkioerrors "cosmossdk.io/errors"
)

// x/burn module sentinel errors
var (
	ErrInvalidSigner = sdkioerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkioerrors.Register(ModuleName, 1101, "sample error")
)

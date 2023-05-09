package app

import (
	"errors"
	"fmt"

	decorators "github.com/BitCannaGlobal/bcna/app/decorators"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

// HandlerOptions extends the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	GovKeeper *govkeeper.Keeper
	Cdc       codec.BinaryCodec
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, fmt.Errorf("%w: account keeper is required for ante builder", errors.New("logic error"))
	}
	if options.BankKeeper == nil {
		return nil, fmt.Errorf("%w: bank keeper is required for ante builder", errors.New("logic error"))
	}
	if options.SignModeHandler == nil {
		return nil, fmt.Errorf("%w:sign mode handler is required for ante builder", errors.New("logic error"))
	}
	if options.GovKeeper == nil {
		return nil, fmt.Errorf("%w:gov keeper is required for AnteHandler", errors.New("logic error"))
	}
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		decorators.NewGovPreventSpamDecorator(options.Cdc, options.GovKeeper),
		// ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}
	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

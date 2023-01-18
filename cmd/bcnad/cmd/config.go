package cmd

import (
	"fmt"

	"github.com/BitCannaGlobal/bcna/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"
	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.SetAddressVerifier(func(bytes []byte) error {
		n := len(bytes)
		if (n != 0) && (n <= address.MaxAddrLen) {
			return nil
		}
		return fmt.Errorf("unexpected address length %d", n)
	})
	config.Seal()
}

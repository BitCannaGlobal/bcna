package app

import (
	"context"
	"fmt"
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"

	storetypes "cosmossdk.io/store/types"

	circuittypes "cosmossdk.io/x/circuit/types"
	"cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	// WASM upgrade/addition
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// RegisterUpgradeHandlers registers upgrade handlers.

func (app App) RegisterUpgradeHandlers() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.StickyFingers(upgradeInfo)
}

func (app *App) StickyFingers(_ upgradetypes.Plan) {
	planName := "stickyfingers"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("Cosmos-SDK v0.50 and WASM are here...")
			// Print the modules with their respective ver.
			for moduleName, version := range fromVM {
				app.Logger().Info(fmt.Sprintf("Module: %s, Version: %d", moduleName, version))

			}
			// Set Consensus Params to avoid err messages
			consensusParams := cmtproto.ConsensusParams{}
			block := cmtproto.BlockParams{
				MaxBytes: 1048576,
				MaxGas:   -1,
			}
			consensusParams.Block = &block
			evidence := cmtproto.EvidenceParams{
				MaxAgeNumBlocks: 100000,
				MaxAgeDuration:  48 * time.Hour,
				MaxBytes:        1048576,
			}
			consensusParams.Evidence = &evidence

			validator := cmtproto.ValidatorParams{
				PubKeyTypes: []string{"ed25519"},
			}
			consensusParams.Validator = &validator

			consensusParams.Version = &cmtproto.VersionParams{}

			err := app.ConsensusParamsKeeper.ParamsStore.Set(ctx, consensusParams)
			app.Logger().Info("Setting the params into the Consensus params keeper...")
			if err != nil {
				app.Logger().Error("Error setting Consensus Params: " + err.Error())
			}

			versionMap, err := app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)

			app.Logger().Info(fmt.Sprintf("post migrate version map: %v", versionMap))

			// Set CosmWasm params
			wasmParams := wasmtypes.DefaultParams()
			// Allowed addresses
			allowedAddresses := []string{
				"bcna1tqywev6xmvrnagfq57c0h5susdy3l789rumufz", // Team 1
				"bcna1csyzlg52g2kd8e0xd6f6elckydhr93ukc3wmqt", // Team 2
				"bcna13jawsn574rf3f0u5rhu7e8n6sayx5gkwgusz73", // Team 3
			}

			// Set CodeUploadAccess to allow Team addresses
			wasmParams.CodeUploadAccess = wasmtypes.AccessConfig{
				Permission: wasmtypes.AccessTypeAnyOfAddresses,
				Addresses:  allowedAddresses,
			}

			// Set InstantiateDefaultPermission to allow Team addresse
			wasmParams.InstantiateDefaultPermission = wasmtypes.AccessTypeEverybody
			err = app.WasmKeeper.SetParams(ctx, wasmParams)
			app.Logger().Info("Setting the params into the Wasm params keeper...")
			if err != nil {
				app.Logger().Error("Error setting the params into the Wasm params keeper...")
				return nil, err
			}
			// return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
			return versionMap, err
		},
	)
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				circuittypes.ModuleName, // commented at v0.50>v0.50 uncomment for v0.47>v0.50
				ibcfeetypes.ModuleName,  // commented at v0.50>v0.50 uncomment for v0.47>v0.50
				nft.ModuleName,          // commented at v0.50>v0.50 uncomment for v0.47>v0.50
				wasmtypes.ModuleName,    // commented at v0.50>v0.50 uncomment for v0.47>v0.50
			},
			Deleted: []string{
				"burn", // commented at v0.50>v0.50 uncomment for v0.47>v0.50
				"bcna", // commented at v0.50>v0.50 uncomment for v0.47>v0.50
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

}

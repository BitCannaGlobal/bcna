package app

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"

	nft "cosmossdk.io/x/nft"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	// Fix Consensus Params
	"github.com/cosmos/cosmos-sdk/runtime"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"

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
			app.Logger().Info("Cosmos-SDK v0.50 and WASM is here...")
			// Print the modules with their respective ver.
			for moduleName, version := range fromVM {
				app.Logger().Info(fmt.Sprintf("Module: %s, Version: %d", moduleName, version))

			}
			// New consensus params keeper using the wrong key again and move the data into the consensus params keeper with the right key
			storesvc := runtime.NewKVStoreService(app.GetKey("upgrade"))
			consensuskeeper := consensuskeeper.NewKeeper(
				app.appCodec,
				storesvc,
				app.AccountKeeper.GetAuthority(),
				runtime.EventService{},
			)

			params, err := consensuskeeper.ParamsStore.Get(ctx)
			app.Logger().Info("Getting the params into the Consensus params keeper...")
			if err != nil {
				app.Logger().Error("Error getting the params into the Consensus params keeper...")
				return nil, err
			}
			err = app.ConsensusParamsKeeper.ParamsStore.Set(ctx, params)
			app.Logger().Info("Setting the params into the Consensus params keeper...")
			if err != nil {
				app.Logger().Error("Error setting the params into the Consensus params keeper...")
				return nil, err
			}
			// Set CosmWasm params
			wasmParams := wasmtypes.DefaultParams()
			wasmParams.CodeUploadAccess = wasmtypes.AllowEverybody //AllowNobody for MainNET
			wasmParams.InstantiateDefaultPermission = wasmtypes.AccessTypeAnyOfAddresses
			err = app.WasmKeeper.SetParams(ctx, wasmParams)
			app.Logger().Info("Setting the params into the Wasm params keeper...")
			if err != nil {
				app.Logger().Error("Error setting the params into the Wasm params keeper...")
				return nil, err
			}
			versionMap, err := app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
			app.Logger().Info(fmt.Sprintf("post migrate version map: %v", versionMap))
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
				wasmtypes.ModuleName,
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

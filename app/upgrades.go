package app

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"

	// nft "cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"
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
	planName := "StickyFingers"
	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("Cosmos-SDK v0.50.x is here...")
			// Print the modules with their respective ver.
			for moduleName, version := range fromVM {
				app.Logger().Info(fmt.Sprintf("Module: %s, Version: %d", moduleName, version))

			}
			versionMap, err := app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
			if err != nil {
				return nil, err
			}
			app.Logger().Info(fmt.Sprintf("post migrate version map: %v", versionMap))
			// return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
			return versionMap, err
		},
	)
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				circuittypes.ModuleName,
				ibcfeetypes.ModuleName,
				// ibctm.ModuleName,
				// solomachine.ModuleName,
				// runtime.ModuleName,
				// "ibc-interchainaccounts",
				// nft.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

}

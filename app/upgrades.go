package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	burnmoduletypes "github.com/BitCannaGlobal/bcna/x/burn/types"
)

var keyTableAssigned = false

// RegisterUpgradeHandlers registers upgrade handlers.

func (app App) RegisterUpgradeHandlers() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.GanjaRevolution47_burn(upgradeInfo)

}

func (app *App) GanjaRevolution47_burn(_ upgradetypes.Plan) {
	planName := "ganjarevolutionburn"
	app.UpgradeKeeper.SetUpgradeHandler(planName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations...adding x/burn module...")
		logger := ctx.Logger().With("upgrade", planName)

		// Run migrations
		logger.Info(fmt.Sprintf("pre migrate version map: %v", fromVM))
		versionMap, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		return versionMap, err
		// return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{

				burnmoduletypes.ModuleName,
				// nft.ModuleName,
			},
		}

		// Configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

// app.IBCKeeper.ClientKeeper

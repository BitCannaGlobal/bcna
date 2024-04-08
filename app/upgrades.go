package app

import (
	"fmt"

	burnmoduletypes "github.com/BitCannaGlobal/bcna/x/burn/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

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

		// Inflation control mechanism - TBD
		// // Get the current params from Mint module
		// mintParams := app.MintKeeper.GetParams(ctx)

		// // Log the params BEFORE apply the new values
		// logger.Info(fmt.Sprintf("Current values for Mint value: InflationMax: %s, InflationMin: %s, InflationRateChange: %s",
		// 	mintParams.InflationMax.String(), mintParams.InflationMin.String(), mintParams.InflationRateChange.String()))

		// // Reduce to half the value of inflation_max, inflation_min and inflation_rate_change
		// mintParams.InflationMax = mintParams.InflationMax.Quo(sdk.NewDec(2))
		// mintParams.InflationMin = mintParams.InflationMin.Quo(sdk.NewDec(2))
		// mintParams.InflationRateChange = mintParams.InflationRateChange.Quo(sdk.NewDec(2))

		// // Set the new values at Mint module
		// if err := app.MintKeeper.SetParams(ctx, mintParams); err != nil {
		// 	return nil, err
		// }

		// // Log the values after apply the changes
		// logger.Info(fmt.Sprintf("New values for Mint value: InflationMax: %s, InflationMin: %s, InflationRateChange: %s",
		// 	mintParams.InflationMax.String(), mintParams.InflationMin.String(), mintParams.InflationRateChange.String()))

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

				burnmoduletypes.ModuleName, // Create the Store for the new module: burn
				// nft.ModuleName,
			},
		}

		// Configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

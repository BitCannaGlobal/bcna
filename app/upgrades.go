package app

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"

	// nft "cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	bcnamoduletypes "github.com/BitCannaGlobal/bcna/x/bcna/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	// "github.com/cosmos/cosmos-sdk/x/nft"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var keyTableAssigned = false

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
	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable

		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true

		// ibc types
		case ibcexported.ModuleName:
			keyTable = icacontrollertypes.ParamKeyTable()
			keyTableAssigned = true
		case ibctransfertypes.ModuleName:
			keyTable = ibctransfertypes.ParamKeyTable()
			keyTableAssigned = true
		case icahosttypes.SubModuleName:
			keyTable = icahosttypes.ParamKeyTable()
			keyTableAssigned = true
		case icacontrollertypes.SubModuleName:
			keyTable = icacontrollertypes.ParamKeyTable()
			keyTableAssigned = true

		// Bitcanna types
		case bcnamoduletypes.ModuleName:
			keyTable = bcnamoduletypes.ParamKeyTable() //nolint:staticcheck
			keyTableAssigned = true

		// Debug:
		default:
			fmt.Println("No matching subspace found:", subspace.Name())
			keyTableAssigned = false
		}

		if !subspace.HasKeyTable() {
			if !keyTableAssigned {
				fmt.Println("KeyTable is not assigned for subspace:", subspace.Name())
			} else {
				subspace.WithKeyTable(keyTable)
			}
		}
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		planName,
		func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("Cosmos-SDK v0.50.x is here...")
			// Print the modules with their respective ver.
			for moduleName, version := range fromVM {
				app.Logger().Info(fmt.Sprintf("Module: %s, Version: %d", moduleName, version))

			}

			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		fmt.Printf("Starting storeUpgrades")
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				circuittypes.ModuleName,
				"feeibc",
				// nft.ModuleName,
			},
		}
		fmt.Printf("Done storeUpgrades: %+v", storeUpgrades)

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

}

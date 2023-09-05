package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	// v047 migration
	"github.com/cosmos/cosmos-sdk/baseapp"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	bcnamoduletypes "github.com/BitCannaGlobal/bcna/x/bcna/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	// "github.com/cosmos/cosmos-sdk/x/nft"
	exported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctmmigrations "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint/migrations"
)

// RegisterUpgradeHandlers registers upgrade handlers.

func (app App) RegisterUpgradeHandlers() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	app.GanjaRevolution47(upgradeInfo)

}

func (app *App) GanjaRevolution47(_ upgradetypes.Plan) {
	planName := "GanjaRevolution"
	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck

		// ibc types
		case ibctransfertypes.ModuleName:
			keyTable = ibctransfertypes.ParamKeyTable()
		case icahosttypes.SubModuleName:
			keyTable = icahosttypes.ParamKeyTable()
		case icacontrollertypes.SubModuleName:
			keyTable = icacontrollertypes.ParamKeyTable()

		// Bitcanna types
		case bcnamoduletypes.ModuleName:
			keyTable = bcnamoduletypes.ParamKeyTable() //nolint:staticcheck
		}

		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(planName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations...")
		logger := ctx.Logger().With("upgrade", planName)

		// Migrate Tendermint consensus parameters from x/params module to a dedicated x/consensus module.
		baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

		// Run migrations
		logger.Info(fmt.Sprintf("pre migrate version map: %v", fromVM))
		versionMap, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		// Note: this migration is optional,
		// You can include x/gov proposal migration documented in [UPGRADING.md](https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md)

		// https://github.com/cosmos/ibc-go/blob/v7.1.0/docs/migrations/v7-to-v7_1.md
		// explicitly update the IBC 02-client params, adding the localhost client type
		params := app.IBCKeeper.ClientKeeper.GetParams(ctx)
		params.AllowedClients = append(params.AllowedClients, exported.Localhost)
		app.IBCKeeper.ClientKeeper.SetParams(ctx, params)

		// OPTIONAL: prune expired tendermint consensus states to save storage space
		if _, err := ibctmmigrations.PruneExpiredConsensusStates(ctx, app.appCodec, app.IBCKeeper.ClientKeeper); err != nil {
			return nil, err
		}

		legacyBaseAppSubspace := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, legacyBaseAppSubspace, &app.ConsensusParamsKeeper)

		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{

				consensustypes.ModuleName,
				crisistypes.ModuleName,
				// nft.ModuleName,
			},
		}

		// Configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

// app.IBCKeeper.ClientKeeper

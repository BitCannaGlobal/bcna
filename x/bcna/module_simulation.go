package bcna

import (
	"math/rand"

	"github.com/BitCannaGlobal/bcna/testutil/sample"
	bcnasimulation "github.com/BitCannaGlobal/bcna/x/bcna/simulation"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = bcnasimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateBitcannaid = "op_weight_msg_bitcannaid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBitcannaid int = 100

	opWeightMsgUpdateBitcannaid = "op_weight_msg_bitcannaid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBitcannaid int = 100

	opWeightMsgDeleteBitcannaid = "op_weight_msg_bitcannaid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBitcannaid int = 100

	opWeightMsgCreateSupplychain = "op_weight_msg_supplychain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSupplychain int = 100

	opWeightMsgUpdateSupplychain = "op_weight_msg_supplychain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSupplychain int = 100

	opWeightMsgDeleteSupplychain = "op_weight_msg_supplychain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSupplychain int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bcnaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		BitcannaidList: []types.Bitcannaid{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		BitcannaidCount: 2,
		SupplychainList: []types.Supplychain{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SupplychainCount: 2,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bcnaGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBitcannaid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBitcannaid, &weightMsgCreateBitcannaid, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBitcannaid = defaultWeightMsgCreateBitcannaid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBitcannaid,
		bcnasimulation.SimulateMsgCreateBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBitcannaid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBitcannaid, &weightMsgUpdateBitcannaid, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBitcannaid = defaultWeightMsgUpdateBitcannaid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBitcannaid,
		bcnasimulation.SimulateMsgUpdateBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBitcannaid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteBitcannaid, &weightMsgDeleteBitcannaid, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBitcannaid = defaultWeightMsgDeleteBitcannaid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBitcannaid,
		bcnasimulation.SimulateMsgDeleteBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSupplychain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSupplychain, &weightMsgCreateSupplychain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSupplychain = defaultWeightMsgCreateSupplychain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSupplychain,
		bcnasimulation.SimulateMsgCreateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSupplychain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSupplychain, &weightMsgUpdateSupplychain, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSupplychain = defaultWeightMsgUpdateSupplychain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSupplychain,
		bcnasimulation.SimulateMsgUpdateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSupplychain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSupplychain, &weightMsgDeleteSupplychain, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSupplychain = defaultWeightMsgDeleteSupplychain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSupplychain,
		bcnasimulation.SimulateMsgDeleteSupplychain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBitcannaid,
			defaultWeightMsgCreateBitcannaid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgCreateBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateBitcannaid,
			defaultWeightMsgUpdateBitcannaid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgUpdateBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteBitcannaid,
			defaultWeightMsgDeleteBitcannaid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgDeleteBitcannaid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSupplychain,
			defaultWeightMsgCreateSupplychain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgCreateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSupplychain,
			defaultWeightMsgUpdateSupplychain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgUpdateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSupplychain,
			defaultWeightMsgDeleteSupplychain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bcnasimulation.SimulateMsgDeleteSupplychain(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
	}
}

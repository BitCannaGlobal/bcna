package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/BitCannaGlobal/bcna/testutil/keeper"
	"github.com/BitCannaGlobal/bcna/testutil/nullify"
	"github.com/BitCannaGlobal/bcna/x/bcna/keeper"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/stretchr/testify/require"
)

func createNBitcannaid(keeper keeper.Keeper, ctx context.Context, n int) []types.Bitcannaid {
	items := make([]types.Bitcannaid, n)
	for i := range items {
		items[i].Id = keeper.AppendBitcannaid(ctx, items[i])
	}
	return items
}

func TestBitcannaidGet(t *testing.T) {
	keeper, ctx := keepertest.BcnaKeeper(t)
	items := createNBitcannaid(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBitcannaid(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBitcannaidRemove(t *testing.T) {
	keeper, ctx := keepertest.BcnaKeeper(t)
	items := createNBitcannaid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBitcannaid(ctx, item.Id)
		_, found := keeper.GetBitcannaid(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBitcannaidGetAll(t *testing.T) {
	keeper, ctx := keepertest.BcnaKeeper(t)
	items := createNBitcannaid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBitcannaid(ctx)),
	)
}

func TestBitcannaidCount(t *testing.T) {
	keeper, ctx := keepertest.BcnaKeeper(t)
	items := createNBitcannaid(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBitcannaidCount(ctx))
}

func TestHasBitcannaidWithBcnaid(t *testing.T) {
	keeper, ctx := keepertest.BcnaKeeper(t)

	// Defines a BitCannaID sample
	exampleBitcannaid := types.Bitcannaid{
		Creator: "creator_address",
		Id:      1,
		Bcnaid:  "test_bcnaid",
	}

	// Saves BitCannaID sample in store
	keeper.SetBitcannaid(ctx, exampleBitcannaid)

	// Test 1: Verify if BitCannaID with the BcnaID of the sample exist
	exist := keeper.HasBitcannaidWithBcnaid(ctx, exampleBitcannaid.Bcnaid)
	require.True(t, exist, "BitCannaID with given BcnaID should exist")

	// Test 2: Verificar si un BitCannaID con un BcnaID diferente no existe
	nonExistingBcnaid := "invented"
	exist = keeper.HasBitcannaidWithBcnaid(ctx, nonExistingBcnaid)
	require.False(t, exist, "BitCannaID with non-existing BcnaID should not exist")
}

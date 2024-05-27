package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/BitCannaGlobal/bcna/testutil/keeper"
	"github.com/BitCannaGlobal/bcna/x/burn/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BurnKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

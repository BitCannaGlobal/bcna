package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
)

func TestBitcannaidMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		concatenated := creator + strconv.Itoa(i) // to produce a different Bcnaid each time
		resp, err := srv.CreateBitcannaid(wctx, &types.MsgCreateBitcannaid{Creator: creator, Bcnaid: concatenated})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestBitcannaidMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBitcannaid
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBitcannaid{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBitcannaid{Creator: "B"},
			err:     types.ErrUnauthorized,
		},
		{
			desc:    "Not found",
			request: &types.MsgUpdateBitcannaid{Creator: creator, Id: 10},
			err:     types.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBitcannaid(ctx, &types.MsgCreateBitcannaid{Creator: creator, Bcnaid: "updated"})
			require.NoError(t, err)

			_, err = srv.UpdateBitcannaid(ctx, tc.request)

			if tc.err != nil {
				require.Contains(t, err.Error(), tc.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBitcannaidMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteBitcannaid
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBitcannaid{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBitcannaid{Creator: "B"},
			err:     types.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBitcannaid{Creator: creator, Id: 10},
			err:     types.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateBitcannaid(wctx, &types.MsgCreateBitcannaid{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBitcannaid(wctx, tc.request)
			if tc.err != nil {
				require.Contains(t, err.Error(), tc.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

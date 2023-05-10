package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
)

func TestBitcannaidMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"

	for i := 0; i < 5; i++ {
		concatenated := strconv.Itoa(i) + creator
		fmt.Println(concatenated)
		resp, err := srv.CreateBitcannaid(ctx, &types.MsgCreateBitcannaid{Creator: creator, Bcnaid: concatenated})
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
			desc:    "Unauthorized",
			request: &types.MsgUpdateBitcannaid{Creator: creator, Id: 10},
			err:     types.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBitcannaid(ctx, &types.MsgCreateBitcannaid{Creator: creator, Bcnaid: "updated"})
			require.NoError(t, err)

			_, err = srv.UpdateBitcannaid(ctx, tc.request)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBitcannaidMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
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
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateBitcannaid(ctx, &types.MsgCreateBitcannaid{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBitcannaid(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

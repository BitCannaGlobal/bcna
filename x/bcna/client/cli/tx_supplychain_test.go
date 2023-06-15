package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/BitCannaGlobal/bcna/testutil/network"
	"github.com/BitCannaGlobal/bcna/x/bcna/client/cli"
)

func TestCreateSupplychain(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz"}
	tests := []struct {
		desc string
		args []string
		err  error
		code uint32
	}{
		{
			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateSupplychain(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
			// } else { RBG: Test here (it comes from https://github.com/RaulBernal/bcna/pull/9/files?file-filters%5B%5D=.go&file-filters%5B%5D=.mod&show-viewed-files=true#diff-785ce63961a15c387203a203fe1d6952bb68c11bdff523ce657b5801459b0abdL46)
			// 	require.NoError(t, err)
			// 	var resp sdk.TxResponse
			// 	require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			// 	require.Equal(t, tc.code, resp.Code)
			// }
		})
	}
}

func TestUpdateSupplychain(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateSupplychain(), args)
	require.NoError(t, err)

	tests := []struct {
		desc string
		id   string
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			id:   "0",
			args: common,
		},
		{
			desc: "key not found",
			id:   "1",
			args: common,
			code: 0x44e,
		},
		{
			desc: "invalid key",
			id:   "invalid",
			err:  strconv.ErrSyntax,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{tc.id}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateSupplychain(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
			// } else {  RBG
			// 	require.NoError(t, err)
			// 	var resp sdk.TxResponse
			// 	require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			// 	require.Equal(t, tc.code, resp.Code)
			// }
		})
	}
}

func TestDeleteSupplychain(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateSupplychain(), args)
	require.NoError(t, err)

	tests := []struct {
		desc string
		id   string
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			id:   "0",
			args: common,
		},
		{
			desc: "key not found",
			id:   "1",
			args: common,
			code: 0x44e,
		},
		{
			desc: "invalid key",
			id:   "invalid",
			err:  strconv.ErrSyntax,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteSupplychain(), append([]string{tc.id}, tc.args...))
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				// } else { RBG
				// 	require.NoError(t, err)
				// 	var resp sdk.TxResponse
				// 	require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				// 	require.Equal(t, tc.code, resp.Code)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

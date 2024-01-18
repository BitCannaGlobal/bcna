package cli_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/BitCannaGlobal/bcna/testutil/network"
	"github.com/BitCannaGlobal/bcna/x/bcna/client/cli"
)

func TestCreateBitcannaid(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	tests := []struct {
		desc   string
		id     string
		fields []string
		args   []string
		err    error
		code   uint32
	}{
		{
			desc: "valid",
			args: common,
		},
		{
			desc:   "BitCannaID already exists",
			id:     "429",
			fields: []string{"xyz", "xyz"},
			args:   common,
			code:   1101,
			// err:  commented then err = 0
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{}
			args = append(args, fields...)
			args = append(args, tc.args...)
			args = append(args, common...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateBitcannaid(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
			// debug RBG
			fmt.Printf("Argumentos: %s\n", args)
			fmt.Printf("Log1: %d\n", tc.code)
			fmt.Printf("Log2: %d\n", resp.Code)
		})
	}
}

func TestUpdateBitcannaid(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz1", "xyz2"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateBitcannaid(), args)
	require.NoError(t, err)

	// tc :=  []struct {
	tests := []struct {
		desc   string
		id     string
		fields []string
		args   []string
		code   uint32
		err    error
	}{
		{
			desc:   "valid",
			id:     "0",
			fields: []string{"Updated1xyz1", "Updated1xyz2"},
			args:   common,
		},
		{
			desc:   "key not found",
			id:     "420",
			fields: []string{"Updated2xyz1", "Updated2xyz2"},
			args:   common,
			code:   0x44e,
			// err:    types.ErrKeyNotFound, commented then err = 0
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{tc.id}
			args = append(args, tc.fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateBitcannaid(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			// debug RBG
			fmt.Printf("Argumentos: %s\n", args)
			fmt.Printf("Log1: %d\n", tc.code)
			fmt.Printf("Log2: %d\n", resp.Code)

			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

func TestDeleteBitcannaid(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateBitcannaid(), args)
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
			code: 0x44e, // sdkerrors.ErrKeyNotFound.ABCICode(), 1102
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteBitcannaid(), append([]string{tc.id}, tc.args...))
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

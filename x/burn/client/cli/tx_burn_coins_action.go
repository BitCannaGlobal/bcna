package cli

import (
	"strconv"

	"github.com/BitCannaGlobal/bcna/x/burn/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurnCoinsAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-coins-action [coins]",
		Short: "Broadcast message BurnCoinsAction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCoins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnCoinsAction(
				clientCtx.GetFromAddress().String(),
				argCoins,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

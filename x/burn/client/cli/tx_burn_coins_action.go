package cli

import (
	"fmt"
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
			argCoins, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse coin: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("failed to get client context: %w", err)
			}

			msg := types.NewMsgBurnCoinsAction(
				clientCtx.GetFromAddress().String(),
				argCoins,
			)
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

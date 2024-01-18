package cli

import (
	"strconv"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListBitcannaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-bitcannaid",
		Short: "list all bitcannaid",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBitcannaidRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BitcannaidAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowBitcannaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-bitcannaid [id]",
		Short: "shows a bitcannaid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetBitcannaidRequest{
				Id: id,
			}

			res, err := queryClient.Bitcannaid(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdSearchBitcannaidByBcnaid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search-bitcannaid [bcnaid]",
		Short: "search a bitcannaid by its bcnaid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			bcnaid := args[0]

			params := &types.QueryGetBitcannaidByBcnaidRequest{
				Bcnaid: bcnaid,
			}

			res, err := queryClient.BitcannaidByBcnaid(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

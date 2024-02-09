package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BitcannaidAll(ctx context.Context, req *types.QueryAllBitcannaidRequest) (*types.QueryAllBitcannaidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bitcannaids []types.Bitcannaid

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bitcannaidStore := prefix.NewStore(store, types.KeyPrefix(types.BitcannaidKey))

	pageRes, err := query.Paginate(bitcannaidStore, req.Pagination, func(key []byte, value []byte) error {
		var bitcannaid types.Bitcannaid
		if err := k.cdc.Unmarshal(value, &bitcannaid); err != nil {
			return err
		}

		bitcannaids = append(bitcannaids, bitcannaid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBitcannaidResponse{Bitcannaid: bitcannaids, Pagination: pageRes}, nil
}

func (k Keeper) Bitcannaid(ctx context.Context, req *types.QueryGetBitcannaidRequest) (*types.QueryGetBitcannaidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	bitcannaid, found := k.GetBitcannaid(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBitcannaidResponse{Bitcannaid: bitcannaid}, nil
}

func (k Keeper) BitcannaidByBcnaid(c context.Context, req *types.QueryGetBitcannaidByBcnaidRequest) (*types.QueryGetBitcannaidByBcnaidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var bitcannaid *types.Bitcannaid
	ctx := sdk.UnwrapSDKContext(c)

	if found := k.HasBitcannaidWithBcnaid(ctx, req.Bcnaid); found {
		bitcannaid, _ = k.GetBitcannaidByBcnaid(ctx, req.Bcnaid)
	} else {
		return nil, status.Error(codes.NotFound, "bitcannaid not found")
	}

	return &types.QueryGetBitcannaidByBcnaidResponse{Bitcannaid: bitcannaid}, nil
}

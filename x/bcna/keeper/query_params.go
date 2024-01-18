package keeper

import (
	"context"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(GoCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(GoCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

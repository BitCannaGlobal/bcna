package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetBitcannaidCount get the total number of bitcannaid
func (k Keeper) GetBitcannaidCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BitcannaidCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBitcannaidCount set the total number of bitcannaid
func (k Keeper) SetBitcannaidCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.BitcannaidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBitcannaid appends a bitcannaid in the store with a new id and update the count
func (k Keeper) AppendBitcannaid(
	ctx context.Context,
	bitcannaid types.Bitcannaid,
) uint64 {
	// Create the bitcannaid
	count := k.GetBitcannaidCount(ctx)

	// Set the ID of the appended value
	bitcannaid.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BitcannaidKey))
	appendedValue := k.cdc.MustMarshal(&bitcannaid)
	store.Set(GetBitcannaidIDBytes(bitcannaid.Id), appendedValue)

	// Update bitcannaid count
	k.SetBitcannaidCount(ctx, count+1)

	return count
}

// SetBitcannaid set a specific bitcannaid in the store
func (k Keeper) SetBitcannaid(ctx context.Context, bitcannaid types.Bitcannaid) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BitcannaidKey))
	b := k.cdc.MustMarshal(&bitcannaid)
	store.Set(GetBitcannaidIDBytes(bitcannaid.Id), b)
}

// GetBitcannaid returns a bitcannaid from its id
func (k Keeper) GetBitcannaid(ctx context.Context, id uint64) (val types.Bitcannaid, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BitcannaidKey))
	b := store.Get(GetBitcannaidIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBitcannaid removes a bitcannaid from the store
func (k Keeper) RemoveBitcannaid(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BitcannaidKey))
	store.Delete(GetBitcannaidIDBytes(id))
}

// GetAllBitcannaid returns all bitcannaid
func (k Keeper) GetAllBitcannaid(ctx context.Context) (list []types.Bitcannaid) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BitcannaidKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bitcannaid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBitcannaidIDBytes returns the byte representation of the ID
func GetBitcannaidIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.BitcannaidKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

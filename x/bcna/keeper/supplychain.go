package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetSupplychainCount get the total number of supplychain
func (k Keeper) GetSupplychainCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SupplychainCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSupplychainCount set the total number of supplychain
func (k Keeper) SetSupplychainCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SupplychainCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSupplychain appends a supplychain in the store with a new id and update the count
func (k Keeper) AppendSupplychain(
	ctx context.Context,
	supplychain types.Supplychain,
) uint64 {
	// Create the supplychain
	count := k.GetSupplychainCount(ctx)

	// Set the ID of the appended value
	supplychain.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SupplychainKey))
	appendedValue := k.cdc.MustMarshal(&supplychain)
	store.Set(GetSupplychainIDBytes(supplychain.Id), appendedValue)

	// Update supplychain count
	k.SetSupplychainCount(ctx, count+1)

	return count
}

// SetSupplychain set a specific supplychain in the store
func (k Keeper) SetSupplychain(ctx context.Context, supplychain types.Supplychain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SupplychainKey))
	b := k.cdc.MustMarshal(&supplychain)
	store.Set(GetSupplychainIDBytes(supplychain.Id), b)
}

// GetSupplychain returns a supplychain from its id
func (k Keeper) GetSupplychain(ctx context.Context, id uint64) (val types.Supplychain, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SupplychainKey))
	b := store.Get(GetSupplychainIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSupplychain removes a supplychain from the store
func (k Keeper) RemoveSupplychain(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SupplychainKey))
	store.Delete(GetSupplychainIDBytes(id))
}

// GetAllSupplychain returns all supplychain
func (k Keeper) GetAllSupplychain(ctx context.Context) (list []types.Supplychain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SupplychainKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Supplychain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSupplychainIDBytes returns the byte representation of the ID
func GetSupplychainIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.SupplychainKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

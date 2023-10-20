package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// GetSupplychainCount get the total number of supplychain
func (k Keeper) GetSupplychainCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
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
func (k Keeper) SetSupplychainCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SupplychainCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSupplychain appends a supplychain in the store with a new id and update the count
func (k Keeper) AppendSupplychain(
	ctx sdk.Context,
	supplychain types.Supplychain,
) uint64 {
	// Create the supplychain
	count := k.GetSupplychainCount(ctx)

	// Set the ID of the appended value
	supplychain.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SupplychainKey))
	appendedValue, err := proto.Marshal(&supplychain)
	if err == nil {
		store.Set(GetSupplychainIDBytes(supplychain.Id), appendedValue)

		// Update supplychain count
		k.SetSupplychainCount(ctx, count+1)

		return count
	} else {
		fmt.Println("DEBUG: err marshaling SupplyChainID")
		return count
	}
}

// SetSupplychain set a specific supplychain in the store
func (k Keeper) SetSupplychain(ctx sdk.Context, supplychain types.Supplychain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SupplychainKey))
	b, err := proto.Marshal(&supplychain)
	if err == nil {
		store.Set(GetSupplychainIDBytes(supplychain.Id), b)
	} else {
		fmt.Println("DEBUG: err setting SupplyChainID")
	}
}

// GetSupplychain returns a supplychain from its id
func (k Keeper) GetSupplychain(ctx sdk.Context, id uint64) (val types.Supplychain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SupplychainKey))
	b := store.Get(GetSupplychainIDBytes(id))
	if b == nil {
		return val, false
	}
	err := proto.Unmarshal(b, &val)
	if err != nil {
		fmt.Println("Error getting the BitCannaID with ID %d: %v\n", id, err)
		return types.Supplychain{}, false
	}
	return val, true
}

// RemoveSupplychain removes a supplychain from the store
func (k Keeper) RemoveSupplychain(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SupplychainKey))
	store.Delete(GetSupplychainIDBytes(id))
}

// GetAllSupplychain returns all supplychain
func (k Keeper) GetAllSupplychain(ctx sdk.Context) (list []types.Supplychain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SupplychainKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Supplychain
		if err := proto.Unmarshal(iterator.Value(), &val); err != nil {
			fmt.Errorf("failed to deserialize SupplyChainID: %w", err)
			continue
		}
		list = append(list, val)
	}

	return
}

// GetSupplychainIDBytes returns the byte representation of the ID
func GetSupplychainIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSupplychainIDFromBytes returns ID in uint64 format from a byte array
func GetSupplychainIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Check if the BitCannaID exist previously
func (k Keeper) HasBitcannaidWithBcnaid(ctx sdk.Context, bcnaid string) bool {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.BitcannaidKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bitcannaid types.Bitcannaid
		proto.Unmarshal(iterator.Value(), &bitcannaid)

		if bitcannaid.Bcnaid == bcnaid {
			return true
		}
	}

	return false
}

func (k Keeper) GetBitcannaidByBcnaid(ctx sdk.Context, bcnaid string) (val *types.Bitcannaid, found bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.BitcannaidKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bitcannaid types.Bitcannaid
		proto.Unmarshal(iterator.Value(), &bitcannaid)

		if bitcannaid.Bcnaid == bcnaid {
			return &bitcannaid, true
		}
	}

	return nil, false
}

// GetBitcannaidCount get the total number of bitcannaid
func (k Keeper) GetBitcannaidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
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
func (k Keeper) SetBitcannaidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BitcannaidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBitcannaid appends a bitcannaid in the store with a new id and update the count
func (k Keeper) AppendBitcannaid(
	ctx sdk.Context,
	bitcannaid types.Bitcannaid,
) uint64 {
	// Create the bitcannaid
	count := k.GetBitcannaidCount(ctx)

	// Set the ID of the appended value
	bitcannaid.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitcannaidKey))
	appendedValue, err := proto.Marshal(&bitcannaid)
	if err == nil {
		store.Set(GetBitcannaidIDBytes(bitcannaid.Id), appendedValue)

		// Update bitcannaid count
		k.SetBitcannaidCount(ctx, count+1)

		return count

	} else {
		fmt.Println("DEBUG: err marshaling BitCannaID")
		return count
	}
}

// SetBitcannaid set a specific bitcannaid in the store
func (k Keeper) SetBitcannaid(ctx sdk.Context, bitcannaid types.Bitcannaid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitcannaidKey))
	b, err := proto.Marshal(&bitcannaid)
	if err == nil {
		store.Set(GetBitcannaidIDBytes(bitcannaid.Id), b)
	} else {
		fmt.Println("DEBUG: err setting BitCannaID")
	}
}

// GetBitcannaid returns a bitcannaid from its id
func (k Keeper) GetBitcannaid(ctx sdk.Context, id uint64) (val types.Bitcannaid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitcannaidKey))
	b := store.Get(GetBitcannaidIDBytes(id))
	if b == nil {
		return val, false
	}
	proto.Unmarshal(b, &val)
	return val, true
}

// RemoveBitcannaid removes a bitcannaid from the store
func (k Keeper) RemoveBitcannaid(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitcannaidKey))
	store.Delete(GetBitcannaidIDBytes(id))
}

// GetAllBitcannaid returns all bitcannaid
func (k Keeper) GetAllBitcannaid(ctx sdk.Context) (list []types.Bitcannaid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BitcannaidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bitcannaid
		proto.Unmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBitcannaidIDBytes returns the byte representation of the ID
func GetBitcannaidIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBitcannaidIDFromBytes returns ID in uint64 format from a byte array
func GetBitcannaidIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

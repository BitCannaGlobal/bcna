package types

const (
	// ModuleName defines the module name
	ModuleName = "bcna"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bcna"
)

var (
	ParamsKey = []byte("p_bcna")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	BitcannaidKey      = "Bitcannaid/value/"
	BitcannaidCountKey = "Bitcannaid/count/"
)

const (
	SupplychainKey      = "Supplychain/value/"
	SupplychainCountKey = "Supplychain/count/"
)

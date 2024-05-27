package types

const (
	// ModuleName defines the module name
	ModuleName = "burn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_burn"
)

var (
	ParamsKey = []byte("p_burn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

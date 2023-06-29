package types

import (
	icqtypes "github.com/strangelove-ventures/async-icq/v7/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "qcrescent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_qcrescent"

	// Version defines the current version the IBC module supports
	Version = icqtypes.Version

	// PortID is the default port id that module binds to
	PortID = "qcrescent"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("qcrescent-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

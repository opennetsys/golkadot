package clientdb

import (
	storagetypes "github.com/opennetsys/go-substrate/client/storage/types"
	types "github.com/opennetsys/go-substrate/types"
)

// StorageFunctionMetadata, StorageFunctionModifier, StorageFunctionType from @polkadot/types/Metadata/Modules
// createFunction from @polkadot/storage/utils/createFunction

// MetadataType ...
type MetadataType struct {
	Key   string
	Value string
}

// SubstrateMetadata ...
type SubstrateMetadata struct {
	Documentation []string
	Type          MetadataType
}

// CreateMethod is small helper function to factorize code on this page
func CreateMethod(method string, key string, meta SubstrateMetadata) types.StorageFunction {
	b := 1
	if meta.Type.Value == "" {
		b = 0
	}
	_ = b

	section := "Block"

	// TODO
	return storagetypes.CreateFunc(
		&section,
		&method,
		&storagetypes.StorageFunctionMetadata{
		/*
			Documentation: NewVector(Text, meta.Documention), // @polkadot/types/codec/Vector
			Modifier: NewStorageFunctionModifier(0),
			Type: NewStorageFunction(meta.Type, b),
		*/
		//ToJSON: func() interface{} {
		//	return key
		//},
		},
		&storagetypes.CreateItemOptions{
			IsUnhashed: false,
			Key:        key,
		},
	)
}

// KeyBestHash ...
func KeyBestHash() types.StorageFunction {
	return CreateMethod("bestHash", "bst:hsh", SubstrateMetadata{
		Documentation: []string{"Best hash"},
		Type: MetadataType{
			Key: "Hash",
		},
	})
}

// KeyBestNumber ...
func KeyBestNumber() types.StorageFunction {
	return CreateMethod("bestNumber", "bst:num", SubstrateMetadata{
		Documentation: []string{"Best block"},
		Type: MetadataType{
			Key: "BlockNumber",
		},
	})
}

// KeyBlockByHash ...
func KeyBlockByHash() types.StorageFunction {
	return CreateMethod("blockByHash", "blk:hsh", SubstrateMetadata{
		Documentation: []string{"Retrieve block by hash"},
		Type: MetadataType{
			Key:   "Hash",
			Value: "Bytes",
		},
	})
}

// KeyHashByNumber ...
func KeyHashByNumber() types.StorageFunction {
	return CreateMethod("hashByNumber", "hsh:num", SubstrateMetadata{
		Documentation: []string{"Retrieve hash by number"},
		Type: MetadataType{
			Key:   "u256",
			Value: "Hash",
		},
	})
}

// KeyHeaderByHash ...
func KeyHeaderByHash() types.StorageFunction {
	return CreateMethod("headerByHash", "hdr:hsh", SubstrateMetadata{
		Documentation: []string{"Retrieve header by hash"},
		Type: MetadataType{
			Key:   "Hash",
			Value: "Bytes",
		},
	})
}

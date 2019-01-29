package db

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

// CreateItemOptions ...
type CreateItemOptions struct {
	IsUnhashed bool
	Key        string
}

// StorageFunctionMetadata ...
type StorageFunctionMetadata struct {
	//Documention *Vector
	//Modifier *Modifier
	//Type *StorageFunction
	//ToJSON *
}

// CreateFunc ...
func CreateFunc(StorageFunctionMetadata, CreateItemOptions) func() {
	// TODO
	return func() {
		// TODO

	}
}

// CreateMethod is small helper function to factorize code on this page
func CreateMethod(method string, key string, meta SubstrateMetadata) func() {
	b := 1
	if meta.Type.Value == "" {
		b = 0
	}
	_ = b

	// TODO
	return CreateFunc(
		//NewText("Block"), // @polkadot/types/Text
		//NewText(method),
		StorageFunctionMetadata{
			/*
				Documentation: NewVector(Text, meta.Documention), // @polkadot/types/codec/Vector
				Modifier: NewStorageFunctionModifier(0),
				Type: NewStorageFunction(meta.Type, b),
			*/
			//ToJSON: func() interface{} {
			//	return key
			//},
		},
		CreateItemOptions{
			IsUnhashed: false,
			Key:        key,
		},
	)
}

// KeyBestHash ...
func KeyBestHash() func() {
	return CreateMethod("bestHash", "bst:hsh", SubstrateMetadata{
		Documentation: []string{"Best hash"},
		Type: MetadataType{
			Key: "Hash",
		},
	})
}

// KeyBestNumber ...
func KeyBestNumber() func() {
	return CreateMethod("bestNumber", "bst:num", SubstrateMetadata{
		Documentation: []string{"Best block"},
		Type: MetadataType{
			Key: "BlockNumber",
		},
	})
}

// KeyBlockByHash ...
func KeyBlockByHash() func() {
	return CreateMethod("blockByHash", "blk:hsh", SubstrateMetadata{
		Documentation: []string{"Retrieve block by hash"},
		Type: MetadataType{
			Key:   "Hash",
			Value: "Bytes",
		},
	})
}

// KeyHashByNumber ...
func KeyHashByNumber() func() {
	return CreateMethod("hashByNumber", "hsh:num", SubstrateMetadata{
		Documentation: []string{"Retrieve hash by number"},
		Type: MetadataType{
			Key:   "u256",
			Value: "Hash",
		},
	})
}

// KeyHeaderByHash ...
func KeyHeaderByHash() func() {
	return CreateMethod("headerByHash", "hdr:hsh", SubstrateMetadata{
		Documentation: []string{"Retrieve header by hash"},
		Type: MetadataType{
			Key:   "Hash",
			Value: "Bytes",
		},
	})
}

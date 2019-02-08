package storagetypes

import (
	types "github.com/opennetsys/golkadot/types"
)

// TODO https://github.com/polkadot-js/api/blob/master/packages/type-storage/src/utils/createFunction.ts#L48

// StorageFunctionMetadata ...
type StorageFunctionMetadata struct {
	//Documention *Vector
	//Modifier *Modifier
	//Type *StorageFunction
	//ToJSON *
}

// CreateItemOptions ...
type CreateItemOptions struct {
	IsUnhashed bool
	Key        string
}

// CreateFunc ...
func CreateFunc(section *string, method *string, meta *StorageFunctionMetadata, options *CreateItemOptions) types.StorageFunction {
	return []uint8{}
}

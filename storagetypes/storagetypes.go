package storagetypes

// TODO ...

// SubstrateType ...
type SubstrateType struct{}

// Substrate ....
var Substrate *SubstrateType

// Code ...
func (s *SubstrateType) Code() []uint8 {
	// TODO
	return nil
}

func init() {
	Substrate = &SubstrateType{}
}

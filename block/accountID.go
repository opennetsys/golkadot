package block

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return EncodeAddress(a)
}

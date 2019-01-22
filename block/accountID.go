package block

// Encode ...
func (a *AccountID) Encode() (string, error) {
	return address.Encode(a[:], nil)
}

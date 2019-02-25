package messages

// Connected ...
type Connected struct {
	Chain string
	Name  string
}

// NewConnected ...
func NewConnected(chain, name string) *Connected {
	return &Connected{
		Chain: chain,
		Name:  name,
	}
}

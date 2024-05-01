package types

type Item struct {
	Name		string
	Quantity	int
}

func NewItem(name string, quantity int) *Item {
	return &Item{
		Name: name,
		Quantity: quantity,
	}
}

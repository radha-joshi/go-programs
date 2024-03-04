package basic

type BasicMenu struct {
	name string
}

func (b *BasicMenu) Name() string {
	return b.name
}

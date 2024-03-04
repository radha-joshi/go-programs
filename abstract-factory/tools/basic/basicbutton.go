package basic

type BasicButton struct {
	name string
}

func (b *BasicButton) Name() string {
	return b.name
}

package basic

type BasicLabel struct {
	name string
}

func (b *BasicLabel) Name() string {
	return b.name
}

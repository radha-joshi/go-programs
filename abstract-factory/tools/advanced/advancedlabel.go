package advanced

type AdvancedLabel struct {
	name string
}

func (b *AdvancedLabel) Name() string {
	return b.name
}

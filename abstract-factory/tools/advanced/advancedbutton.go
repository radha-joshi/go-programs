package advanced

type AdvancedButton struct {
	name string
}

func (b *AdvancedButton) Name() string {
	return b.name
}

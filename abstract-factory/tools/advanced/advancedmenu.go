package advanced

type AdvancedMenu struct {
	name string
}

func (b *AdvancedMenu) Name() string {
	return b.name
}

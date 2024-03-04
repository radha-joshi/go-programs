package advanced

import "abstractfactory/tools"

type AdvancedToolFactory struct {
}

func (f *AdvancedToolFactory) CreateButton(buttonType string) tools.Button {
	return &AdvancedButton{
		name: "AdvancedButton",
	}
}

func (f *AdvancedToolFactory) CreateMenu(menuType string) tools.Menu {
	return &AdvancedMenu{
		name: "AdvancedButton",
	}
}

func (f *AdvancedToolFactory) CreateLabel(labelType string) tools.Label {
	return &AdvancedLabel{
		name: "AdvancedButton",
	}
}

func New() tools.ToolFactory {
	return &AdvancedToolFactory{}
}

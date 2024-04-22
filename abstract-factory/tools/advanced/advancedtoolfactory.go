package advanced

import "abstractfactory/tools"

type advancedToolFactory struct {
}

func (f *advancedToolFactory) CreateButton(buttonType string) tools.Button {
	return &AdvancedButton{
		name: "AdvancedButton",
	}
}

func (f *advancedToolFactory) CreateMenu(menuType string) tools.Menu {
	return &AdvancedMenu{
		name: "AdvancedButton",
	}
}

func (f *advancedToolFactory) CreateLabel(labelType string) tools.Label {
	return &AdvancedLabel{
		name: "AdvancedButton",
	}
}

func New() tools.ToolFactory {
	return &advancedToolFactory{}
}

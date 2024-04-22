package basic

import "abstractfactory/tools"

type basicToolFactory struct {
}

func (f *basicToolFactory) CreateButton(buttonType string) tools.Button {
	return &BasicButton{
		name: "BasicButton",
	}
}

func (f *basicToolFactory) CreateMenu(menuType string) tools.Menu {
	return &BasicMenu{
		name: "BasicMenu",
	}
}

func (f *basicToolFactory) CreateLabel(labelType string) tools.Label {
	return &BasicLabel{
		name: "BasicLabel",
	}
}

func New() tools.ToolFactory {
	return &basicToolFactory{}
}

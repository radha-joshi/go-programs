package basic

import "abstractfactory/tools"

type BasicToolFactory struct {
}

func (f *BasicToolFactory) CreateButton(buttonType string) tools.Button {
	return &BasicButton{
		name: "BasicButton",
	}
}

func (f *BasicToolFactory) CreateMenu(menuType string) tools.Menu {
	return &BasicMenu{
		name: "BasicButton",
	}
}

func (f *BasicToolFactory) CreateLabel(labelType string) tools.Label {
	return &BasicLabel{
		name: "BasicButton",
	}
}

func New() tools.ToolFactory {
	return &BasicToolFactory{}
}

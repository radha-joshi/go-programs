// Package basic implements the ToolFactory interface from the abstractfactory/tools package.
// It provides basic implementations for UI components like buttons, menus, and labels.
package basic

import "abstractfactory/tools"

// basicToolFactory is a concrete factory that produces basic UI components.
// It implements the ToolFactory interface defined in the abstractfactory/tools package.
type basicToolFactory struct{}

// CreateButton creates a basic button component.
// It ignores the buttonType argument and always returns a BasicButton with a fixed name.
func (f *basicToolFactory) CreateButton(buttonType string) tools.Button {
	return &BasicButton{
		name: "BasicButton",
	}
}

// CreateMenu creates a basic menu component.
// It ignores the menuType argument and always returns a BasicMenu with a fixed name.
func (f *basicToolFactory) CreateMenu(menuType string) tools.Menu {
	return &BasicMenu{
		name: "BasicMenu",
	}
}

// CreateLabel creates a basic label component.
// It ignores the labelType argument and always returns a BasicLabel with a fixed name.
func (f *basicToolFactory) CreateLabel(labelType string) tools.Label {
	return &BasicLabel{
		name: "BasicLabel",
	}
}

// New returns a new instance of a basicToolFactory.
// This function allows the creation of a factory that can produce basic UI components.
func New() tools.ToolFactory {
	return &basicToolFactory{}
}

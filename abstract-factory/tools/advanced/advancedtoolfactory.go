// Package advanced implements the ToolFactory interface from the abstractfactory/tools package.
// It provides advanced implementations for UI components like buttons, menus, and labels.
package advanced

import "abstractfactory/tools"

// advancedToolFactory is a concrete factory that produces advanced UI components.
// It implements the ToolFactory interface defined in the abstractfactory/tools package.
type advancedToolFactory struct{}

// CreateButton creates an advanced button component.
// It ignores the buttonType argument and always returns an AdvancedButton with a fixed name.
func (f *advancedToolFactory) CreateButton(buttonType string) tools.Button {
	return &AdvancedButton{
		name: "AdvancedButton",
	}
}

// CreateMenu creates an advanced menu component.
// It ignores the menuType argument and always returns an AdvancedMenu with a fixed name.
func (f *advancedToolFactory) CreateMenu(menuType string) tools.Menu {
	return &AdvancedMenu{
		name: "AdvancedMenu",
	}
}

// CreateLabel creates an advanced label component.
// It ignores the labelType argument and always returns an AdvancedLabel with a fixed name.
func (f *advancedToolFactory) CreateLabel(labelType string) tools.Label {
	return &AdvancedLabel{
		name: "AdvancedLabel",
	}
}

// New returns a new instance of an advancedToolFactory.
// This function facilitates the creation of a factory that can produce advanced UI components.
func New() tools.ToolFactory {
	return &advancedToolFactory{}
}

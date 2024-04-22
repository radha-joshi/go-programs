// Package tools defines an interface for a factory that creates UI components.
// This interface is designed to be implemented by concrete factories that produce specific types of UI components,
// such as buttons, menus, and labels.
package tools

// ToolFactory is an interface that defines methods for creating UI components.
// Implementers of this interface are expected to provide concrete instances of buttons, menus, and labels.
type ToolFactory interface {
	// CreateButton creates a Button based on the specified buttonType.
	// The buttonType parameter allows for differentiation between different styles or behaviors of buttons.
	CreateButton(buttonType string) Button

	// CreateMenu creates a Menu based on the specified menuType.
	// The menuType parameter allows for differentiation between different styles or configurations of menus.
	CreateMenu(menuType string) Menu

	// CreateLabel creates a Label based on the specified labelType.
	// The labelType parameter allows for differentiation between different styles or usages of labels.
	CreateLabel(labelType string) Label
}

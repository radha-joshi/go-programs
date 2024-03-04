package tools

type ToolFactory interface {
	CreateButton(buttonType string) Button
	CreateMenu(menuType string) Menu
	CreateLabel(labelType string) Label
}

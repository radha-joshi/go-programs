// Package main demonstrates the usage of abstract factories for creating UI components.
// It utilizes both 'basic' and 'advanced' implementations from the abstractfactory/tools package.
package main

import (
	"abstractfactory/tools"
	"abstractfactory/tools/advanced"
	"abstractfactory/tools/basic"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Application represents the main structure for this example application.
type Application struct{}

// Start runs the application by utilizing two different factories to create UI components.
// It demonstrates the creation of UI components using both basic and advanced factories.
func (a *Application) Start() {
	useFactory(basic.New())    // Use the basic tool factory to create basic style components.
	useFactory(advanced.New()) // Use the advanced tool factory to create advanced style components.
}

// useFactory demonstrates the creation of different UI components using a given ToolFactory.
// It creates a button, label, and menu using the factory provided and logs their names.
func useFactory(factory tools.ToolFactory) {
	button := factory.CreateButton("someType")               // Create a button of "someType", though type handling is not implemented.
	log.Info().Msg(fmt.Sprintf("Button: %s", button.Name())) // Log the name of the button.

	label := factory.CreateLabel("someType")               // Create a label of "someType", similar to button.
	log.Info().Msg(fmt.Sprintf("Label: %s", label.Name())) // Log the name of the label.

	menu := factory.CreateMenu("someType")               // Create a menu of "someType", similar to button and label.
	log.Info().Msg(fmt.Sprintf("Menu: %s", menu.Name())) // Log the name of the menu.
}

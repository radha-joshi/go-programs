package main

import (
	"abstractfactory/tools"
	"abstractfactory/tools/advanced"
	"abstractfactory/tools/basic"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	useFactory(basic.New())
	useFactory(advanced.New())
}

func useFactory(factory tools.ToolFactory) {

	button := factory.CreateButton("someType")
	log.Info().Msg(fmt.Sprintf("Button: %s", button.Name()))

	label := factory.CreateLabel("someType")
	log.Info().Msg(fmt.Sprintf("Label: %s", label.Name()))

	menu := factory.CreateMenu("someType")
	log.Info().Msg(fmt.Sprintf("Menu: %s", menu.Name()))

}

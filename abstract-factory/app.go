package main

import (
	"abstractfactory/tools"
	"abstractfactory/tools/basic"

	"github.com/rs/zerolog/log"
)

func start() {
	var factory tools.ToolFactory = basic.New()
	button := factory.CreateButton("someType")

	log.Info().Msg("Button: ", button.Name())

}

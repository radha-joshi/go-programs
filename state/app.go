package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

type Display struct {
	opacity bool
}

func (d *Display) show(message string) {
	if d.opacity {
		log.Info().Msg(fmt.Sprintf("With opacity: %s", message))
	} else {
		log.Info().Msg(fmt.Sprintf("Without opacity: %s", message))
	}
}

func (d *Display) setOpacity(value bool) {
	d.opacity = value
}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	myDisplay := Display{}
	myDisplay.setOpacity(true)
	myDisplay.show("Show Message")
	myDisplay.setOpacity(false)
	myDisplay.show("Show Message Again")
}

package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type StateApplication struct{}

type SuperDisplay struct {
	opacity bool
	state   DisplayState
}

type DisplayState interface {
	show(message string)
}

type OpaqueState struct{}

func (s *OpaqueState) show(message string) {
	log.Info().Msg(fmt.Sprintf("From State With opacity: %s", message))
}

type TransparentState struct{}

func (s *TransparentState) show(message string) {
	log.Info().Msg(fmt.Sprintf("From State With opacity: %s", message))
}

func (d *SuperDisplay) show(message string) {
	if d.state != nil {
		d.state.show(message)
	}
}

func (d *SuperDisplay) setOpacity(value bool) {
	d.opacity = value
	if value {
		d.state = &OpaqueState{}
	} else {
		d.state = &TransparentState{}
	}
}

func (a *StateApplication) Start() {
	log.Info().Msg("Application started.")

	myDisplay := SuperDisplay{}
	myDisplay.setOpacity(true)
	myDisplay.show("Show Message")
	myDisplay.setOpacity(false)
	myDisplay.show("Show Message Again")
}

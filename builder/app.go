package main

import (
	"fmt"

	"company.com/builder/person"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	builder := person.PersonBuilder{}
	person := builder.
		WithName("Acharya").
		WithCity("Nagpur").
		WithDesignation("Teacher").
		Build()
	log.Info().Msg(fmt.Sprintf("Person: %v", person))
}

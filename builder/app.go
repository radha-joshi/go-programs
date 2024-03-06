package main

import (
	"fmt"

	"company.com/builder/person"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	builder := person.NewBuilder()
	person1 := builder.
		WithName("Acharya").
		WithCity("Nagpur").
		WithDesignation("Teacher").
		Build()
	log.Info().Msg(fmt.Sprintf("Person: %v", person1))

	person1 = person.NewBuilder().
		WithCity("Raipur").
		WithName("Bachchan").
		Build()
	log.Info().Msg(fmt.Sprintf("Person: %v", person1))
}

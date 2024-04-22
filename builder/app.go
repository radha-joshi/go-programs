// Package main demonstrates the usage of the builder pattern for creating person objects.
// It uses the person builder from the "company.com/builder/person" package to configure and create person entities.
package main

import (
	"fmt"

	"company.com/builder/person"
	"github.com/rs/zerolog/log"
)

// Application is a struct that serves as the base for the application demonstration.
type Application struct{}

// Start runs the application, demonstrating the use of the PersonBuilder to create and log person entities.
func (a *Application) Start() {
	// Initialize a new PersonBuilder from the person package.
	builder := person.NewBuilder()
	// Build a person with specified name, city, and designation.
	person1 := builder.
		WithName("Acharya").
		WithCity("Nagpur").
		WithDesignation("Teacher").
		Build()
	// Log the details of the created person.
	log.Info().Msg(fmt.Sprintf("Person: %v", person1))

	// Build another person using a new builder, specifying different attributes.
	person1 = person.NewBuilder().
		WithCity("Raipur").
		WithName("Bachchan").
		Build()
	// Log the details of the second created person.
	log.Info().Msg(fmt.Sprintf("Person: %v", person1))

	city := ""

	person2Builder := person.NewBuilder()
	if city != "" {
		person2Builder.WithCity(city)
	}
	person2 := person2Builder.
		WithName("Bachchan").
		Build()
	log.Info().Msg(fmt.Sprintf("Person: %v", person2))
}

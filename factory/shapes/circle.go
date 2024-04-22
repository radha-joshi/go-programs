// Package shapes provides definitions and functionalities for various geometric shapes.
// It currently includes a concrete implementation for a circle.
package shapes

import "github.com/rs/zerolog/log"

// Circle represents a geometric shape with a specific name.
type Circle struct {
	name string // name is the private field that holds the name of the shape.
}

// Name returns the name of the circle.
// This method satisfies the Shape interface requirement.
func (c *Circle) Name() string {
	return c.name
}

// NewCircle creates a new instance of a Circle, logs the creation, and returns it as a Shape type.
// This function demonstrates logging during the creation of shape instances.
func NewCircle() Shape {
	log.Info().Msg("New Circle created.")
	return &Circle{
		name: "Circle",
	}
}

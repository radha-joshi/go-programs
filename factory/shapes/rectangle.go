// Package shapes provides definitions and functionalities for various geometric shapes.
// It includes a concrete implementation for a rectangle.
package shapes

import "github.com/rs/zerolog/log"

// Rectangle represents a geometric shape with a specific name.
type Rectangle struct {
	name string // name is the private field that holds the name of the shape.
}

// Name returns the name of the rectangle.
// This method satisfies the Shape interface requirement.
func (c *Rectangle) Name() string {
	return c.name
}

// NewRectangle creates a new instance of a Rectangle, logs the creation, and returns it as a Shape type.
// This function demonstrates logging during the creation of shape instances.
func NewRectangle() Shape {
	log.Info().Msg("New Rectangle created.")
	return &Rectangle{
		name: "Rectangle",
	}
}

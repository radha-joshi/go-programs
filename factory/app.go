// Package main demonstrates the usage of the shapes package, specifically utilizing
// a factory pattern to create different types of shapes.
package main

import (
	"factory/shapes"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Application struct serves as a base structure for the application.
type Application struct{}

// Start initializes the application, creates shapes using a factory,
// and demonstrates processing and logging shape information.
func (a *Application) Start() {
	// Create a new shape factory from the shapes package.
	shapeFactory := shapes.NewShapeFactory()

	// Create a circle and rectangle using the shape factory.
	circle, _ := shapeFactory.CreateShape("circle")
	rect, _ := shapeFactory.CreateShape("rectangle")

	// Print the names and types of the created shapes.
	fmt.Printf("Shape name: %s, Type: %T\n", circle.Name(), circle)
	fmt.Printf("Shape name: %s, Type: %T\n", rect.Name(), rect)

	// Initialize a slice to store shapes.
	shapeList := make([]shapes.Shape, 5)

	// Create three circles and add them to the shape list.
	for i := 0; i < 3; i++ {
		if shp, error := shapeFactory.CreateShape("circle"); error == nil {
			shapeList[i] = shp
		}
	}

	// Create two rectangles and add them to the shape list.
	for i := 3; i < 5; i++ {
		if shp, error := shapeFactory.CreateShape("rectangle"); error == nil {
			shapeList[i] = shp
		}
	}

	// Log the names of the shapes in the list.
	for i := 0; i < 5; i++ {
		log.Info().Msg(fmt.Sprintf("Shape name: %s", shapeList[i].Name()))
	}
}

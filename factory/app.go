package main

import (
	"factory/shapes"
	"fmt"

	"github.com/rs/zerolog/log"
)

func start() {
	circle, _ := shapes.CreateShape("circle")
	rect, _ := shapes.CreateShape("rectangle")

	fmt.Printf("Shape name: %s, Type: %T\n", circle.Name(), circle)
	fmt.Printf("Shape name: %s, Type: %T\n", rect.Name(), rect)

	shapeList := make([]shapes.Shape, 5)
	for i := 0; i < 3; i++ {
		if shp, error := shapes.CreateShape("circle"); error == nil {
			shapeList[i] = shp
		}
	}
	for i := 3; i < 5; i++ {
		if shp, error := shapes.CreateShape("rectangle"); error == nil {
			shapeList[i] = shp
		}
	}

	for i := 0; i < 5; i++ {
		log.Info().Msg(fmt.Sprintf("Shape name: %s", shapeList[i].Name()))
	}

}

package main

import (
	"factory/shapes"
	"fmt"
)

func start() {
	circle, _ := shapes.CreateShape("circle")
	rect, _ := shapes.CreateShape("rectangle")

	fmt.Printf("Shape name: %s, Type: %T\n", circle.Name(), circle)
	fmt.Printf("Shape name: %s, Type: %T\n", rect.Name(), rect)
}

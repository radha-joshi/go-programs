package shapes

import "fmt"

func CreateShape(shapeType string) (Shape, error) {
	if shapeType == "circle" {
		return NewCircle(), nil
	}
	if shapeType == "rectangle" {
		return NewRectangle(), nil
	}
	return nil, fmt.Errorf("Wrong shape type passed")
}

// Package shapes provides a factory for creating different types of shapes.
// It supports creation of predefined shapes like circles and rectangles.
package shapes

import "fmt"

// ShapeFactory is an interface that defines a method to create shapes based on a type identifier.
type ShapeFactory interface {
	// CreateShape creates a Shape object based on the specified shapeType.
	// It returns a Shape and an error. The error is non-nil if the shapeType is not recognized.
	CreateShape(shapeType string) (Shape, error)
}

// shapeFactory is a private struct that implements the ShapeFactory interface.
type shapeFactory struct{}

// CreateShape returns a Shape based on the shapeType argument.
// It supports creating "circle" and "rectangle". If an unsupported shapeType is provided,
// it returns an error indicating an unrecognized shape type.
func (sf *shapeFactory) CreateShape(shapeType string) (Shape, error) {
	switch shapeType {
	case "circle":
		return NewCircle(), nil
	case "rectangle":
		return NewRectangle(), nil
	default:
		return nil, fmt.Errorf("wrong shape type passed: %s", shapeType)
	}
}

// NewShapeFactory creates and returns a new instance of a ShapeFactory.
// This function provides access to the concrete implementation of the factory.
func NewShapeFactory() ShapeFactory {
	return &shapeFactory{}
}

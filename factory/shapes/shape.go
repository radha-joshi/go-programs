// Package shapes defines an interface for shape entities within the application.
// It provides a foundational structure for shapes, enabling polymorphic behavior.
package shapes

// Shape defines the basic interface for all shape types.
// It requires implementing a method that returns the name of the shape.
type Shape interface {
	// Name returns the name of the shape as a string.
	Name() string
}

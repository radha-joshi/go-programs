// Package person defines structures and interfaces to create and manage person entities.
package person

// PersonBuilder is an interface for building Person objects step-by-step.
// It provides methods to set the name, city, and designation of a person,
// and to build the final Person object.
type PersonBuilder interface {
	// WithName sets the name of the person and returns the updated PersonBuilder.
	WithName(name string) PersonBuilder

	// WithCity sets the city of the person and returns the updated PersonBuilder.
	WithCity(city string) PersonBuilder

	// WithDesignation sets the designation of the person and returns the updated PersonBuilder.
	WithDesignation(designation string) PersonBuilder

	// Build constructs the final Person object incorporating all provided details.
	Build() Person
}

// personBuilder implements the PersonBuilder interface.
// It holds temporary state for building a Person object.
type personBuilder struct {
	name        string
	city        string
	designation string
}

// WithName assigns the name to the person being built and returns the PersonBuilder instance.
func (b *personBuilder) WithName(name string) PersonBuilder {
	b.name = name
	return b
}

// WithCity assigns the city to the person being built and returns the PersonBuilder instance.
func (b *personBuilder) WithCity(city string) PersonBuilder {
	b.city = city
	return b
}

// WithDesignation assigns the designation to the person being built and returns the PersonBuilder instance.
func (b *personBuilder) WithDesignation(designation string) PersonBuilder {
	b.designation = designation
	return b
}

// NewBuilder creates and returns a new instance of a PersonBuilder.
func NewBuilder() PersonBuilder {
	return &personBuilder{}
}

// Build constructs a new Person object using the details stored in the builder.
func (b *personBuilder) Build() Person {
	return Person{
		name:        b.name,
		city:        b.city,
		designation: b.designation,
	}
}

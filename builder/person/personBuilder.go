package person

type PersonBuilder interface {
	WithName(name string) PersonBuilder
	WithCity(city string) PersonBuilder
	WithDesignation(designation string) PersonBuilder
	Build() Person
}

type personBuilder struct {
	name        string
	city        string
	designation string
}

func (b *personBuilder) WithName(name string) PersonBuilder {
	b.name = name
	return b
}

func (b *personBuilder) WithCity(city string) PersonBuilder {
	b.city = city
	return b
}

func (b *personBuilder) WithDesignation(designation string) PersonBuilder {
	b.designation = designation
	return b
}

func NewBuilder() PersonBuilder {
	return &personBuilder{}
}

func (b *personBuilder) Build() Person {
	return Person{
		name:        b.name,
		city:        b.city,
		designation: b.designation,
	}
}

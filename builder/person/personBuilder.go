package person

type PersonBuilder struct {
	name        string
	city        string
	designation string
}

func (b *PersonBuilder) WithName(name string) *PersonBuilder {
	b.name = name
	return b
}

func (b *PersonBuilder) WithCity(city string) *PersonBuilder {
	b.city = city
	return b
}

func (b *PersonBuilder) WithDesignation(designation string) *PersonBuilder {
	b.designation = designation
	return b
}

func (b *PersonBuilder) Build() Person {
	return Person{
		name:        b.name,
		city:        b.city,
		designation: b.designation,
	}
}

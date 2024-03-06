package person

type Person struct {
	name        string
	city        string
	designation string
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetCity() string {
	return p.city
}

func (p *Person) SetCity(city string) {
	p.city = city
}

func (p *Person) GetDesignation() string {
	return p.designation
}

func (p *Person) SetDesignation(designation string) {
	p.designation = designation
}

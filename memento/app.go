package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type State interface {
	Data() string
}

type ProductState struct {
	data string
}

func (s *ProductState) Data() string {
	return s.data
}

type Product struct {
	complexData string
}

func (p *Product) GetState() State {
	return &ProductState{
		data: p.complexData,
	}
}

func NewProduct(state State) Product {
	return Product{
		complexData: state.Data(),
	}
}

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	// first run
	product1 := NewProduct(&ProductState{
		data: "complex data",
	})
	fmt.Println(product1)

	savedState := product1.GetState()
	// system exit

	// system started

	// second run
	product2 := NewProduct(savedState)
	fmt.Println(product2)
}

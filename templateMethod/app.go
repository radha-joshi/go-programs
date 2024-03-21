package main

import (
	"fmt"

	"company.com/templateMethod/insurance"
	"github.com/rs/zerolog/log"
)

type Application struct{}

type MyDiscountStrategy struct{}

// template method implemented by client
func (s *MyDiscountStrategy) GetDiscount() int {
	return 15
}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	calculator := insurance.Calculator{
		DiscountStrategy: &MyDiscountStrategy{},
	}
	log.Info().Msg(fmt.Sprintf("With My Discount %v", calculator.GetFinalAmount(100)))
}

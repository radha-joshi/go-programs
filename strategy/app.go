package main

import (
	"fmt"

	"company.com/strategy/insurance"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	calculator := insurance.Calculator{
		DiscountStrategy: &insurance.SimpleDiscount{},
	}
	log.Info().Msg(fmt.Sprintf("With Simple Discount %v", calculator.GetFinalAmount(100)))

	calculator = insurance.Calculator{
		DiscountStrategy: &insurance.ComplexDiscount{},
	}
	log.Info().Msg(fmt.Sprintf("With Complex Discount %v", calculator.GetFinalAmount(100)))
}

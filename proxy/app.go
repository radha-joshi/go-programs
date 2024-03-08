package main

import (
	"github.com/rs/zerolog/log"

	"company.com/proxy/paymentGateway"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	var paymentService paymentGateway.PaymentService = paymentGateway.GetPaymentService()
	paymentService.ProcessPayment()
}

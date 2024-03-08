package paymentGateway

import (
	"company.com/proxy/paypal"
	"github.com/rs/zerolog/log"
)

type PaymentService interface {
	ProcessPayment()
}

type paymentGateway struct {
}

func (p *paymentGateway) ProcessPayment() {
	log.Info().Msg("Delegating to Paypal library")
	var provider *paypal.Provider = paypal.GetProvider()
	provider.Process()
}

func GetPaymentService() PaymentService {
	return &paymentGateway{}
}

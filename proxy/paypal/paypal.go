package paypal

import "github.com/rs/zerolog/log"

type Provider struct {
}

func (p *Provider) Process() {
	log.Info().Msg("Paypal provider processing payment")
}

func GetProvider() *Provider {
	return &Provider{}
}

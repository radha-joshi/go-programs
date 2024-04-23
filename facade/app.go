package main

import (
	"company.com/facade/banking"
	"company.com/facade/finance"
	"company.com/facade/locker"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	var account banking.Account = banking.Account{}
	var box locker.Box = locker.Box{}

	var bank banking.Bank = banking.Bank{}
	var locker locker.Locker = locker.Locker{}

	var financeOps *finance.FinStorageFacade

	financeOps = finance.NewFinStorageFacade(&bank)

	financeOps.StoreAmount(&account, 23.8)

	financeOps = finance.NewFinStorageFacade(&locker)

	financeOps.StoreAmount(&box, 123.8)

}

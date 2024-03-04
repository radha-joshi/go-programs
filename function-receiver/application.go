package main

import (
	"fmt"

	"company.com/functionreceiver/util"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) start() {
	log.Info().Msg("It is working")

	myUtils := util.NewMyUtils()
	log.Info().Msg(fmt.Sprintf("Type of myUtils %T", myUtils))
	log.Info().Msg(myUtils.Reverse("Some Text"))

	superUtils := util.NewMySuperUtils()
	log.Info().Msg(fmt.Sprintf("Type of mySuperUtils %T", superUtils))
	log.Info().Msg(superUtils.Reverse("Some Text"))
}

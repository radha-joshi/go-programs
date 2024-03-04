package main

import (
	"fmt"

	"company.com/singletontry/util"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) start() {
	log.Info().Msg("It is working")

	myUtils := util.NewMyUtils()
	myUtils.SetName("One")
	log.Info().Msg(fmt.Sprintf("Type of myUtils %T", myUtils))
	log.Info().Msg(myUtils.Reverse("Some Text"))

	myUtils2 := util.NewMyUtils()
	myUtils2.SetName("two")
	log.Info().Msg(fmt.Sprintf("Type of myUtils %T", myUtils2))
	log.Info().Msg(myUtils2.Reverse("Some Text"))

	log.Info().Msg(fmt.Sprintf("Are objects same: %v", myUtils == myUtils2))
	log.Info().Msg(fmt.Sprintf("myUtils name: %s", myUtils.GetName()))
	log.Info().Msg(fmt.Sprintf("myUtils2 name: %s", myUtils2.GetName()))
}

package main

import (
	"fmt"

	"company.com/iterator/calendar"
	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	var calendar calendar.MonthList = *calendar.NewMonthList()
	for calendar.HasNext() {
		log.Info().Msg(fmt.Sprintf("Month: %v", calendar.Next()))
	}

	var custom calendar.CustomList = *calendar.NewCustomList()
	for custom.HasNext() {
		log.Info().Msg(fmt.Sprintf("Month: %v", custom.Next()))
	}
}

package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func display(text string) {
	log.Info().Msg(text)
}

func displayWithQuotes(text string) {
	log.Info().Msg(fmt.Sprintf("\"%s\"\n", text))
}

type Function func(string)

func logDecorator(f Function) Function {
	return func(text string) {
		log.Info().Msg("Before calling function.")
		f(text)
		log.Info().Msg("After calling function")
	}
}

func lineDecorator(f Function) Function {
	return func(text string) {
		log.Info().Msg("-------------------------------")
		f(text)
		log.Info().Msg("-------------------------------")
	}
}

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")
	display("Some Text")
	displayWithQuotes("Some Text")

	logDisplay := logDecorator(display)
	logDisplay("Some Text Again")

	logDisplayWithQuotes := logDecorator(displayWithQuotes)
	logDisplayWithQuotes("Some Text Again")

	displayWithDecorators := lineDecorator(logDecorator(display))
	displayWithDecorators("Checkin the decorators")
}

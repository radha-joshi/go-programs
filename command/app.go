package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"company.com/command/processor"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	// implement command stack to enable undo operation
	// o _ _
	// _ _ x
	// _ _ _

	// NextMove(x, y, state)

	//1) 1000
	//2) (- 2) 998
	//3) (- 10) 988 <- undo

	var processor *processor.Processor = processor.NewProcessor()
	processor.AddValue(10) // 10
	processor.AddValue(20) // 30
	processor.AddValue(30) // 60
	log.Info().Msg(fmt.Sprintf("Value is now: %v", processor.GetValue()))

	processor.Undo()
	log.Info().Msg(fmt.Sprintf("Value is now: %v", processor.GetValue()))

	processor.Undo()
	log.Info().Msg(fmt.Sprintf("Value is now: %v", processor.GetValue()))
}

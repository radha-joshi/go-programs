package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	editor := Editor{}

	editor.AddText("AAAAAAA", "BOLD", 10, "RED")
	editor.AddText("BBBBB", "BOLD", 10, "RED")
	editor.AddText("CCCCCC", "NORMAL", 10, "RED")
	editor.AddText("DDDDD", "BOLD", 10, "RED")
	editor.AddText("eEEEEE", "BOLD", 10, "RED")
	editor.AddText("FFFFF", "BOLD", 11, "RED")
	editor.AddText("GGGGG", "BOLD", 10, "BLUE")
	editor.AddText("HHHHH", "BOLD", 10, "RED")

	fmt.Println(&editor)
}

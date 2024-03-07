package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	folder1 := NewFolder("Folder1")
	folder1.AddComponent(NewFile("File1"))
	folder1.AddComponent(NewFile("File2"))

	folder2 := NewFolder("Folder2")
	folder1.AddComponent(NewFile("File3"))
	folder2.AddComponent(folder1)
	fmt.Println(folder2)

}

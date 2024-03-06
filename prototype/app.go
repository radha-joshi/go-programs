package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

func (a *Application) Start() {
	folder1 := NewFolder("F1")

	folder2 := folder1.Clone()

	log.Info().Msg(fmt.Sprintf("is Folder2 same as folder1 %v", folder2 == folder1))
	log.Info().Msg(fmt.Sprintf("Folder2 name %s, folder1 name %s", folder2.GetName(), folder1.GetName()))

	myFolder := NewFolder("MyFolder")
	myFile1 := NewFile("File1")
	myFolder.AddFile(myFile1)

	log.Info().Msg(fmt.Sprintf("Folder %v", myFolder))

	copyFolder := myFolder.Clone()

	myFile1.SetName("NewName")

	log.Info().Msg(fmt.Sprintf("myFolder %v", myFolder))
	log.Info().Msg(fmt.Sprintf("copyFolder %v", copyFolder))
}

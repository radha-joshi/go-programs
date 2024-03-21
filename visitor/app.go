package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Application struct{}

type MyVisitor struct{}

func (v *MyVisitor) Visit(component Component) {
	log.Info().Msg(fmt.Sprintf("Visiting component: %v-%v", component.GetName(), component.GetType()))
}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	visitor := &MyVisitor{}

	folder1 := NewFolder("Folder1")
	folder1.AddComponent(NewFile("File1"))
	folder1.AddComponent(NewFile("File2"))

	folder2 := NewFolder("Folder2")
	folder1.AddComponent(NewFile("File3"))
	folder2.AddComponent(folder1)

	folder2.Accept(visitor)
}

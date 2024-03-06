package main

import "fmt"

type Node interface {
	GetName() string
	SetName(name string)
	Clone() Node
}

type Folder struct {
	name     string
	fileList []*File
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) SetName(name string) {
	f.name = name
}

func (f *Folder) AddFile(file *File) {
	f.fileList = append(f.fileList, file)
}

func (f *Folder) String() string {
	value := f.GetName()
	for _, file := range f.fileList {
		value += fmt.Sprintf("%s\n", file.GetName())
	}
	return value
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		fileList: make([]*File, 0),
	}
}

func (f *Folder) Clone() Node {
	folder := &Folder{}
	folder.name = f.name
	for _, file := range f.fileList {
		folder.AddFile((file.Clone()).(*File))
	}
	return folder
}

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) SetName(name string) {
	f.name = name
}

func (f *File) Clone() Node {
	return &File{
		name: f.GetName(),
	}
}

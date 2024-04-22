// Package main demonstrates a composite pattern and cloning functionality for file system nodes.
// It defines interfaces and types that represent folders and files within a file system.
package main

import "fmt"

// Node defines the interface for components in the file system (both files and folders).
// It requires methods for handling names and cloning the nodes.
type Node interface {
	GetName() string     // GetName returns the name of the node.
	SetName(name string) // SetName sets the name of the node.
	Clone() Node         // Clone creates a deep copy of the node.
}

// Folder represents a folder in a file system, containing other folders and files.
type Folder struct {
	name     string  // name is the name of the folder.
	fileList []*File // fileList contains all files in the folder.
}

// GetName returns the name of the folder.
func (f *Folder) GetName() string {
	return f.name
}

// SetName sets the name of the folder.
func (f *Folder) SetName(name string) {
	f.name = name
}

// AddFile adds a file to the folder.
func (f *Folder) AddFile(file *File) {
	f.fileList = append(f.fileList, file)
}

// String returns a string representation of the folder and its contents.
func (f *Folder) String() string {
	value := f.GetName()
	for _, file := range f.fileList {
		value += fmt.Sprintf("%s\n", file.GetName())
	}
	return value
}

// NewFolder creates and returns a new Folder with the given name.
func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		fileList: make([]*File, 0),
	}
}

// Clone creates a deep copy of the Folder, including all contained Files.
func (f *Folder) Clone() Node {
	folder := &Folder{}
	folder.name = f.name
	for _, file := range f.fileList {
		folder.AddFile((file.Clone()).(*File))
	}
	return folder
}

// File represents a file in a file system.
type File struct {
	name string // name is the name of the file.
}

// NewFile creates and returns a new File with the specified name.
func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

// GetName returns the name of the file.
func (f *File) GetName() string {
	return f.name
}

// SetName sets the name of the file.
func (f *File) SetName(name string) {
	f.name = name
}

// Clone creates a deep copy of the File.
func (f *File) Clone() Node {
	return &File{
		name: f.GetName(),
	}
}

package main

import "fmt"

type Component interface {
	GetType() string
	GetName() string
	GetParent() Component
	GetChildren() []Component
	AddComponent(component Component)
	SetParent(parent Component)
}

type BaseComponent struct {
	componentType string
	name          string
	parent        Component
	children      []Component
}

func (c *BaseComponent) GetType() string {
	return c.componentType
}

func (c *BaseComponent) GetName() string {
	return c.name
}

func (c *BaseComponent) GetParent() Component {
	return c.parent
}

func (c *BaseComponent) SetParent(parent Component) {
	c.parent = parent
}

func (c *BaseComponent) GetChildren() []Component {
	return c.children
}

func (c *BaseComponent) AddComponent(component Component) {
	component.SetParent(c)
	c.children = append(c.children, component)
}

type File struct {
	BaseComponent
}

type Folder struct {
	BaseComponent
}

func NewFile(name string) *File {
	return &File{
		BaseComponent{
			name:          name,
			componentType: "File",
			parent:        nil,
			children:      nil,
		},
	}
}

func (f *File) String() string {
	return fmt.Sprintf("Type: %s, Name: %s", f.componentType, f.name)
}

func NewFolder(name string) *Folder {
	return &Folder{
		BaseComponent{
			name:          name,
			componentType: "Folder",
			parent:        nil,
			children:      make([]Component, 0),
		},
	}
}

func (f *Folder) String() string {
	value := fmt.Sprintf("Type: %s, Name: %s", f.componentType, f.name)
	value += "\nChildren:"
	for _, item := range f.children {
		value += fmt.Sprintf("\nParent: %v, %v", item.GetParent().GetName(), item)
	}
	return value
}

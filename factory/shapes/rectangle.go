package shapes

import "github.com/rs/zerolog/log"

type Rectangle struct {
	name string
}

func (c *Rectangle) Name() string {
	return c.name
}

func NewRectangle() Shape {
	log.Info().Msg("New Rectangle created.")
	return &Rectangle{
		name: "Rectangle",
	}
}

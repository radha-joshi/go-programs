package shapes

import "github.com/rs/zerolog/log"

type Circle struct {
	name string
}

func (c *Circle) Name() string {
	return c.name
}

func NewCircle() Shape {
	log.Info().Msg("New Circle created.")
	return &Circle{
		name: "Circle",
	}
}

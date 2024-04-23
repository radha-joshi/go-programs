package main

import (
	"fmt"

	"company.com/adapter/vertexes"
	"github.com/rs/zerolog/log"
)

type Application struct{}

type Coordinate interface {
	X() float32
	Y() float32
}

type coordinate struct {
	x float32
	y float32
}

func (c *coordinate) X() float32 {
	return c.x
}

func (c *coordinate) Y() float32 {
	return c.y
}

func NewCoordinate(x float32, y float32) Coordinate {
	return &coordinate{
		x: x,
		y: y,
	}
}

type Polyline struct {
	pointList []Coordinate
}

func NewPolyline() *Polyline {
	return &Polyline{
		pointList: make([]Coordinate, 0),
	}
}

func (p *Polyline) AddCoordinate(c Coordinate) {
	p.pointList = append(p.pointList, c)
}

type PolygonAdaptor struct {
	polyline *Polyline
}

func (a *PolygonAdaptor) GetPoint(index int64) vertexes.Point {
	return a.polyline.pointList[index]
}

func (a *PolygonAdaptor) GetCount() int {
	return len(a.polyline.pointList)
}

func NewPolygonAdaptor(polyline *Polyline) vertexes.Polygon {
	return &PolygonAdaptor{
		polyline: polyline,
	}
}

func (a *Application) Start() {
	log.Info().Msg("Application started.")

	p := NewPolyline()
	p.AddCoordinate(NewCoordinate(1.0, 2.0))
	p.AddCoordinate(NewCoordinate(2.0, 5.0))
	p.AddCoordinate(NewCoordinate(4.0, 8.0))

	polygonAdapter := NewPolygonAdaptor(p)

	distance := vertexes.CalculateDistance(polygonAdapter)
	log.Info().Msg(fmt.Sprintf("Distance is %v", distance))
}

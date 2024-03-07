package main

import (
	"company.com/adapter/vertexes"
	"github.com/rs/zerolog/log"
)

type Application struct{}

type Coordinate interface {
	X() float32
	Y() float32
}

type Polyline struct {
	pointList []Coordinate
}

func NewPolyline() *Polyline {
	return &Polyline{
		pointList: make([]Coordinate, 0),
	}
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

	polygonAdapter := NewPolygonAdaptor(p)

	vertexes.CalculateDistances(polygonAdapter)
}

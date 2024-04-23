package vertexes

import "github.com/rs/zerolog/log"

type Point interface {
	X() float32
	Y() float32
}

type Polygon interface {
	GetPoint(index int64) Point
	GetCount() int
}

type myPolygon struct {
	pointList []Point
}

func (p *myPolygon) GetPoint(index int64) Point {
	return p.pointList[index]
}

func (p *myPolygon) GetCount() int {
	return len(p.pointList)
}

func NewPolygon() Polygon {
	return &myPolygon{
		pointList: make([]Point, 0),
	}
}

func CalculateDistance(polygon Polygon) int64 {
	log.Info().Msg("calculateDistance worked")
	// Actual logic will be complicated to calculate distance
	// This is sample return value
	return int64(polygon.GetCount()) * 100
}

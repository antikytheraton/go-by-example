package geometry

import "math"

// Point represents a point in space
type Point struct{ X, Y float64 }

// Distance function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy has a pointer receiver tp Point struct
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

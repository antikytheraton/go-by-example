package geometry

import "image/color"

// ColoredPoint is a composed struct
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// // Distance is a wrapper function
// func (p ColoredPoint) Distance(q Point) float64 {
// 	return p.Point.Distance(q)
// }
//
// // ScaleBy is a wrapper function with a pointer receiver
// func (p ColoredPoint) ScaleBy(factor float64) {
// 	p.Point.ScaleBy(factor)
// }

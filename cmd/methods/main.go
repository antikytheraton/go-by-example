package main

import (
	"fmt"
	"image/color"

	"github.com/antikytheraton/go-context/cmd/methods/geometry"
)

func main() {

	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	fmt.Printf("function call: %f\n", geometry.Distance(p, q))
	fmt.Printf("method call: %f\n", p.Distance(q))

	perim := geometry.Path{
		geometry.Point{X: 1, Y: 1},
		geometry.Point{X: 5, Y: 1},
		geometry.Point{X: 5, Y: 4},
		geometry.Point{X: 1, Y: 1},
	}
	fmt.Printf("Perim: %f\n", perim.Distance())

	fmt.Printf("p: %f\n", geometry.Point{X: 1, Y: 2}.Distance(q))

	// Pointers
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Printf("pptr: %f\n", p)

	// Composing Types by struct embeding
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Printf("x: %f\n", cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Printf("y: %f\n", cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var ccp = geometry.ColoredPoint{Point: geometry.Point{X: 1, Y: 1}, Color: red}
	var ccq = geometry.ColoredPoint{Point: geometry.Point{X: 5, Y: 4}, Color: blue}
	fmt.Printf("compose distance: %f\n", ccp.Distance(ccq.Point)) // "5"
	ccp.ScaleBy(2)
	ccq.ScaleBy(2)
	fmt.Printf("compose distance: %f\n", ccp.Distance(ccq.Point))

	// Method values and expressions

	vp := geometry.Point{X: 1, Y: 2}
	vq := geometry.Point{X: 4, Y: 6}

	distanceFromP := vp.Distance                               // method value
	fmt.Printf("method value: %f\n", distanceFromP(vq))        // "5"
	var origin geometry.Point                                  // {0, 0}
	fmt.Printf("origin distance: %f\n", distanceFromP(origin)) // "2.2360" sqrt(5)

	scaleP := vp.ScaleBy // method value
	scaleP(2)            // vp becomes (2,4)
	scaleP(3)            //       then (6, 12)
	scaleP(10)           //       then (60, 120)
	fmt.Printf("scaleP: %f\n", vp)

	distance := geometry.Point.Distance
	fmt.Printf("distance: %f\n", distance(vp, vq))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&vp, 2)
	fmt.Printf("scale: %f\n", vp)
	fmt.Printf("%T\n", scale)

}

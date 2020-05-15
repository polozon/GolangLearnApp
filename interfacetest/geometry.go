package interfacetest

import (
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}
func (r rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func GetIt(typeOfThing int) Geometry {
	if typeOfThing == 1 {
		return rect{width: 3, height: 4}
	}
	return circle{radius: 5}
}

package basics

import (
	"errors"
	"fmt"
	"math"
)

/* Interfaces - collection of method signatures */
type geometry interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

/* Method implementations */
func (cr circle) area() float64 {
	fmt.Print("area of radius ", cr.radius, " is ")
	return math.Pi * cr.radius * cr.radius
}

func (rec rectangle) area() float64 {
	fmt.Print("area of rectangle width ", rec.width, " height ", rec.height, " is ")
	return rec.width * rec.height
}

func area(geo geometry) float64 {
	return geo.area()
}

/* Embedding of structs/interfaces */
type circleext struct {
	circle // Embedding using field without name
	x, y   int
}

func divide(n1 float64, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("cannot divide by zero") // Return with an error
	}
	return n1 / n2, nil // Return with nil error
}

func Basics6() {
	cr := circle{radius: 5}
	rec := rectangle{width: 2, height: 3}

	// Using interface based implementation
	fmt.Println(area(cr))
	fmt.Println(area(rec))

	// Using embedded struct
	extcr := circleext{circle: circle{radius: 1}, x: 2, y: 3}
	fmt.Println(extcr, "radius is", extcr.circle.radius)
	fmt.Println(area(extcr.circle))

	// Error handling
	result, err := divide(10, 0)
	if err == nil {
		fmt.Println("result is", result)
	} else {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10
	fmt.Println("area = ", c1.getArea())
}
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

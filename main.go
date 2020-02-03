package main

import (
	"fmt"
	"math"
)

// Circle struct
type Circle struct {
	Radius int
}

// Area circle
func (c Circle) Area() float64 {
	fRad := float64(c.Radius)
	return math.Pi * fRad * fRad
}

func main() {

	fmt.Println(Circle{2}.Area())
	fmt.Println(Circle{3}.Area())
	fmt.Println(Circle{4}.Area())
	fmt.Println(Circle{5}.Area())
	fmt.Println(Circle{6}.Area())
}

package interfaces

import (
	"fmt"
	"math"
)

type Speaker interface {
	Speak()
}

type Dog struct{}

func (Dog) Speak() {
	fmt.Println("Woof!")
}

func SimpleInterface() {
	s := Dog{}
	s.Speak()
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64      { return s.Side * s.Side }
func (s Square) Perimeter() float64 { return 4 * s.Side }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintAllShapes(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("Shape is of type: %T, area: %.2f, perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}

func MultipleMethodsImplementation() {
	PrintAllShapes(
		Circle{Radius: 5},
		Square{5},
		Circle{3},
		Square{Side: 10},
		Circle{Radius: 30},
	)
}

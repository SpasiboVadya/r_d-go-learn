package methods

import (
	"fmt"
	"lesson04/interfaces"
)

type Point struct {
	X, Y int
}

func (p Point) MoveBy(dx, dy int) {
	p.X += dx
	p.Y += dy
	fmt.Printf("Inside value receiver: %+v\n", p)
}

func (p *Point) Shift(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func ValueVsPointer() {
	p := Point{1, 2}
	fmt.Printf("Initial point: %+v\n", p)
	p.MoveBy(2, 2)
	fmt.Printf("Point after MoveBy: %+v\n", p)
	p.Shift(2, 2)
	fmt.Printf("Point after Shift: %+v\n", p)
}

type MyInt int

func NewMyInt(x int) MyInt {
	return MyInt(x)
}

func (x MyInt) IsPositive() bool {
	return x > 0
}

func CallCustomTypeMethod() {
	x := MyInt(42)
	fmt.Printf("number %d is positive: %v", x, x.IsPositive())
}

func (p Point) GetString() interfaces.Point {
	return interfaces.Point{A: p.X, B: p.Y}
}

func MethodValueAndExpression() {
	point := Point{1, 2}
	methodValue := point.GetString
	fmt.Printf("Point method value: %+v\n", methodValue())

	methodExpression := Point.GetString
	fmt.Printf("Point method expression: %+v", methodExpression(point))
}

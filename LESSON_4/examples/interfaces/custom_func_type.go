package interfaces

import "fmt"

type Adder interface {
	Add(a, b int) int
}

type addFunc func(int, int) int

func (f addFunc) Add(a, b int) int {
	return f(a, b)
}

func ShowFuncInterface() {
	var a Adder = addFunc(func(x, y int) int { return x + y })
	fmt.Println(a.Add(3, 4))
}

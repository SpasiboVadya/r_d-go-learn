package empty_interface

import "fmt"

func describe(value any) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func BasicExample() {
	var x interface{}
	x = 42
	fmt.Println(x)

	x = "hello"
	fmt.Println(x)

	x = []int{1, 2, 3}
	fmt.Println(x)

	describe(3.14)
	describe("text")
	describe(struct {
		X int
		Y string
	}{X: 10, Y: "hello"})
}

func SliceOfAny() {
	items := []any{123, "hello", true, 3.14}
	for _, item := range items {
		fmt.Printf("Item: %v, Type: %T\n", item, item)
	}
}

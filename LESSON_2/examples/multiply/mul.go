package main

import (
	"fmt"
	"strconv"
)

func add(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func createUser() error {
	return nil
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type is int:", v)
	case string:
		fmt.Println("Type is string:", v)
	case bool:
		fmt.Println("Type is bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func ifNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else if num%3 == 0 {
		fmt.Println(num, "is divisible by 3")
	} else {
		fmt.Println(num, "is odd")
	}
}

func switchNum(num int) {
	switch {
	case num%2 == 0:
		fmt.Println(num, "is even")
	case num%3 == 0:
		fmt.Println(num, "is divisible by 3")
	default:
		fmt.Println(num, "is odd")
	}
}

func setNameValue(name string) {
	name = "Gopher"
}

func setName(name *string) {
	n := "Goph"
	name = &n
	fmt.Println(*name)
	//*name = "Gopher"
}

func main() {
	if strconv.Itoa(1) == "1" && 2 < 4 {
		println("hello")
	}
}

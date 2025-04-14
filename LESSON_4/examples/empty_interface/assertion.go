package empty_interface

import "fmt"

func printType(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case func() bool:
		fmt.Println("func:", v())
	default:
		fmt.Println("unknown type")
	}
}

func TypeAssertions() {
	var val any = "Gophers"

	str, ok := val.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}

	printType(val)
	printType(10)
	printType(func() bool { return false })
}

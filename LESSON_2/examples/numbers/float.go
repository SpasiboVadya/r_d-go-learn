package main

import "fmt"

const unit = 1e8 // 100_000_000

var a int64 = 12345678901 // $123.45678901
var b int64 = 99          // $0.00000099

func swap(x, y string) (string, string) {

	return y, x
}

func main() {
	//a := 0.1
	//b := 0.2
	//sum := a + b
	//
	//fmt.Println("a + b =", sum)
	//fmt.Println("a + b == 0.3:", sum == 0.3)

	w, _ := swap("hello", "world")
	fmt.Println(w)
}

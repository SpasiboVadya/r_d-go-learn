package main

import (
	"fmt"
)

func PrintArr(arr [4]int) {
	fmt.Println(&arr[0])
	fmt.Println(arr)
}

func main() {
	arr := [...]int{1, 2, 3, 4}
	p0 := &arr[0]
	p1 := &arr[1]
	p2 := &arr[2]

	fmt.Println(p0)
	fmt.Println(p1)
	fmt.Println(p2)

	PrintArr(arr)
}

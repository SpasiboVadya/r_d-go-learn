package main

import (
	"fmt"
)

func PrintSlice(arr []int16) {
	var ptr *int16
	if len(arr) > 0 {
		ptr = &arr[0]
	}
	fmt.Printf("cap: %d, len: %d, addr: %p\n", cap(arr), len(arr), ptr)
}

func addElement(arr []int) {
	arr = append(arr, 1)
}

func main() {
	sls := []int{1, 2, 3} //1, 2, 3 len 3 cap 3
	sls = append(sls, 4)  // 1, 2, 3, 4, len 4 cap 6
	addElement(sls)
	fmt.Println(sls[0:6])
	// unsafe.SliceData
	// reflect.SliceHeader

	//var arr []int16
	//PrintSlice(arr)
	////
	//arr = append(arr, 1)
	////
	//PrintSlice(arr)
	//
	//for i := 0; i < 1000; i++ {
	//	arr = append(arr, 1)
	//	PrintSlice(arr)
	//}

	//PrintSlice(arr[1:])
	//AddBalance()
	// DON'T FORGET TO SHOW example with slice on slice
}

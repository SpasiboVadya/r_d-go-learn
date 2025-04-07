package main

import "fmt"

var cache = map[string]string{}

func do(a string) string {
	// some heavy calculations ....
	return a + a
}

func doWithCache(a string) string {
	if val, ok := cache[a]; ok {
		fmt.Printf("%s from cache\n", val)
		return val
	}

	result := do(a)
	cache[a] = result
	return result
}

func main() {
	AddBalance()
	//fmt.Println(doWithCache("a"))
	//fmt.Println(doWithCache("b"))
	//fmt.Println(doWithCache("c"))
	//fmt.Println(doWithCache("a"))
	//
	//fmt.Println("First loop")
	//for k, v := range cache {
	//	fmt.Println(k, v)
	//}
	//fmt.Println("Second loop")
	//for k, v := range cache {
	//	fmt.Println(k, v)
	//}
	//fmt.Println("Third loop")
	//fmt.Println(doWithCache("d"))
	//for k, v := range cache {
	//	fmt.Println(k, v)
	//}
	//fmt.Printf("First print %+v\n", cache)
	//fmt.Printf("Second print %+v\n", cache)
}

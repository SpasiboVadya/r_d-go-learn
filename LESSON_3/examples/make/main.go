package main

import (
	"fmt"
	"math/big"
	"runtime"
)

func printMem(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %.2f MB\n", label, float64(m.Alloc)/1024/1024)
}

func main() {
	b := new(big.Int).Add(big.NewInt(1), big.NewInt(1))
	fmt.Println(b)
	//sliceGrowth()
	//mapGrowth()
}

package main

import "fmt"

// Benchmark_sliceGrowth-12    	    2694	    399555 ns/op
// Benchmark_sliceGrowth-12    	    2869	    405667 ns/op
// Benchmark_sliceGrowth-12    	     512	   2375082 ns/op
func sliceGrowth() {
	size := 1_000_000
	var m []int = make([]int, 0, size)
	m = append(m, 2)
	m[0] = 1
	printMem(fmt.Sprintf("After %d inserts", 0))
	for i := 0; i < size; i++ {
		m = append(m, i)

		if i%100_000 == 0 {
			printMem(fmt.Sprintf("After %d inserts", i))
		}
	}
}

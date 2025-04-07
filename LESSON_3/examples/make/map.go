package main

import (
	"strconv"
)

// Benchmark_mapGrowth-12    	      12	  89519194 ns/op
// Benchmark_mapGrowth-12    	       9	 116077519 ns/op

func mapGrowth() {
	size := 1_000_000
	m := make(map[string]int)

	for i := 0; i <= size; i++ {
		m[strconv.Itoa(i)] = i

		//if i%100_000 == 0 {
		//	printMem(fmt.Sprintf("After %d inserts", i))
		//}
	}
}

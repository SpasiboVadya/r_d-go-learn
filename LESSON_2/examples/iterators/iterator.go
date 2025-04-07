package main

import (
	"fmt"
	"iter"
)

func CountTo(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func MyMapIter(m map[string]int) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func main() {
	for v := range CountTo(3) {
		fmt.Println(v)
	}

	//for k, v := range MyMapIter(map[string]int{"a": 1, "b": 2}) {
	//	fmt.Println(k, v)
	//}

	pipeline := Take(
		Map(
			Filter(
				CountFrom(2),
				func(v int) bool { return v%2 == 0 },
			),
			func(v int) int { return v * v },
		),
		5,
	)

	for v := range pipeline {
		fmt.Println(v)
	}
}

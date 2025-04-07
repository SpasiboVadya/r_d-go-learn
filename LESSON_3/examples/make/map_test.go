package main

import "testing"

func Benchmark_mapGrowth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapGrowth()
	}
}

func Benchmark_sliceGrowth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceGrowth()
	}
}

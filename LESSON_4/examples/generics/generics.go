package generics

import "fmt"

func Sum[R ~int | float64 | float32](nums []R) R {
	var total R
	for _, n := range nums {
		total += n
	}
	return total
}

type MyInt int

type MyFloat float64

func SumGenerics() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]MyInt{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5}))
	//fmt.Println(Sum([]MyFloat{1.5, 2.5}))
}

func Map[T any, R any](in []T, fn func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func MapGenerics() {
	squares := Map([]int{1, 2, 3}, func(n int) int {
		return n * n
	})
	fmt.Println(squares)
}

type Pair[T any, U any] struct {
	First  T
	Second U
}

func PairGenerics() {
	p := Pair[int, string]{First: 1, Second: "one"}
	fmt.Println(p)
}

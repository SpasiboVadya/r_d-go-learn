package main

import (
	"iter"
)

func CountFrom(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Filter(seq iter.Seq[int], pred func(int) bool) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Map(seq iter.Seq[int], m func(int) int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := range seq {
			if !yield(m(v)) {
				return
			}
		}
	}
}

func Take(seq iter.Seq[int], n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		count := 0
		for v := range seq {
			if count >= n {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}

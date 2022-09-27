package main

import (
	"fmt"
	"math/rand"
)

type customStack struct {
	stack [][]int
}

func WRandint(min int, max int) int {

	return min + rand.Intn(max-min)
}

func Append(size int, value int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		t := WRandint(1, size+1)
		for j := 0; j < t; j++ {
			m[i] = append(m[i], value)
		}
		if len(m) > 1 {
			for h := 0; h < len(m); h++ {
				if len(m[i]) < 1 {
					m = append(m[:h], m[h+1:]...)
					h--
				}
			}

			for r := 0; r < 1; r++ {
				m = append(m[i:], m[:i]...)
			}
		}
	}
	return m
}

func merge(a [][]int, b [][]int) [][]int {
	final := make([][]int, 1)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if size(a[i]) < size(b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(items [][]int) [][]int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func size(a []int) int {
	return len(a)
}

func main() {
	var n int
	fmt.Scanln(&n)
	unsorted := Append(n, 1)
	fmt.Println(unsorted)
	sorted := mergeSort(unsorted)
	for i := 0; i < len(sorted); i++ {
		if len(sorted[i]) < 1 {
			sorted = append(sorted[:i], sorted[i+1:]...)
			i--
		}
	}
	fmt.Println(sorted)

}

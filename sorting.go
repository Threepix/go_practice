package main

import (
	"fmt"
	"math/rand"
)

func selectionSort(n int) {
	posled := make([]int, n)
	poss := make([]*int, n)
	for i := 0; i < n; i++ {
		posled[i] = rand.Intn(100)
		poss[i] = &posled[i]
	}
	len_poss := len(poss)
	fmt.Println("unsorted")
	for i := 0; i < len_poss; i++ {
		fmt.Println(*poss[i])
	}
	for i := 0; i < len_poss; i++ {
		var minInd = i
		for j := i; j < len_poss; j++ {
			if *poss[j] < *poss[minInd] {
				minInd = j
			}
		}
		poss[i], poss[minInd] = poss[minInd], poss[i]
	}
	fmt.Println("sorted")
	for i := 0; i < len_poss; i++ {
		fmt.Println(*poss[i])
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	selectionSort(n)
}

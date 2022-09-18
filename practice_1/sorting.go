package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(n int) {
	count := 0
	now := time.Now()

	posled := make([]int, n) //n
	poss := make([]*int, n)  //n
	for i := 0; i < n; i++ { //2n
		posled[i] = rand.Intn(100)
		poss[i] = &posled[i]
	}
	len_poss := len(poss) //n
	fmt.Println("unsorted")
	for i := 0; i < len_poss; i++ { //n
		fmt.Println(*poss[i])
	}
	for i := 0; i < len_poss; i++ { // 4n
		var minInd = i
		for j := i; j < len_poss; j++ { //2n
			if *poss[j] < *poss[minInd] {
				minInd = j
			}
		}
		poss[i], poss[minInd] = poss[minInd], poss[i]
	}
	fmt.Println("sorted")           //n
	for i := 0; i < len_poss; i++ { //n
		fmt.Println(*poss[i])
	}
	durat := time.Since(now)
	fmt.Println(durat) //n
	count = n + n + n*(2*n) + n + n*n + n*4*n + n + n*n + 3*n
	fmt.Println(count) //n
}

func main() {
	var n int
	fmt.Scanln(&n)
	selectionSort(n)
}


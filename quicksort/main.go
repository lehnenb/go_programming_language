package main

import (
	"fmt"
	"math/rand"
	"time"
)

var interactions int = 0

func main() {
	rand.Seed(time.Now().Unix())
	//sli := rand.Perm(100)
	sli := []int{8, 9, 1, 3, 10, 18, 5}
	sort(sli, 0, len(sli)-1)
	fmt.Printf("number of iterations: %d", interactions)
	fmt.Printf("ordered list %v", sli)
}

func sort(sli []int, start int, end int) {
	if start < end {
		pIndex := partition(sli, start, end)
		sort(sli, start, pIndex)
		sort(sli, pIndex+1, end)
	}
}

func partition(sli []int, start int, end int) int {
	fmt.Print("\n\n\n")
	pivot := sli[end]
	pIndex := start - 1

	for index := start; index < end; index++ {
		interactions++
		val := sli[index]
		fmt.Printf("index = %d | pIndex %d | val = %d | pivot = %d\n", index, pIndex, val, pivot)

		if val < pivot {
			// fmt.Printf("increased lowest element from %d to %d ", index-1, index)
			// fmt.Println("")
			// fmt.Printf("lowest element val = %d | current value = %d", sli[pIndex], sli[index])
			// fmt.Println("")
			pIndex++
			sli[index], sli[pIndex] = sli[pIndex], sli[index]
			// fmt.Printf("lowest element val = %d | current value = %d", sli[pIndex], sli[index])
			// fmt.Println("")
		}
	}

	// fmt.Printf("biggest element val = %d | pivot value = %d\n", sli[end], sli[pIndex+1])
	sli[end], sli[pIndex+1] = sli[pIndex+1], sli[end]
	fmt.Printf("Final pIndex %d, List: %v", pIndex, sli)
	// fmt.Printf("biggest element val = %d | pivot value = %d", sli[end], sli[pIndex+1])
	// fmt.Println("")
	// fmt.Printf("array: %v", sli)
	// fmt.Print("\n\n\n")

	return pIndex
}

/*
	Store initial value of the partition
	i = 0

	Main comparison loop - Responsible for swapping elements and updating the partition point:
		j = 0
		[8, 9, 3, 4, 5]

		i = 0
		j = 1
		[8, 9, 3, 4, 5]

		i = 0
		j = 2
		[8, 9, 3, 4, 5]
		expl: since "i" is smaller than the pivot, swap current element with current parition point and update partition point


		i = 1
		j = 3
		[3, 9, 8, 4, 5]

		i = 2
		j = 4
		[3, 4, 8, 9, 5]

	Final swap, responsible for putting pivot in its right index:
	(since we haven't touched it because its comparing with < instead of <=)
		i := 2
*/

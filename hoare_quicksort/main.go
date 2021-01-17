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

	sli := []int{22, 2, 1, 18}
	sort(sli, 0, len(sli)-1)

	fmt.Printf("%v", sli)
}

func sort(sli []int, start int, end int) {
	if start < end {
		pIndex := partition(sli, start, end)
		fmt.Printf("Pivot: %d \n", pIndex)
		sort(sli, start, pIndex)
		sort(sli, pIndex+1, end)
	}
}

func partition(sli []int, start int, end int) int {
	fmt.Print("\n\n")
	fmt.Printf("Start: %d - End: %d \n", start, end)
	fmt.Printf("Slice: %v \n", sli[start:end+1])
	pivot := sli[(start+end)/2]
	fmt.Printf("Chosen pivot\n index: %d \n value: %d \n", (start+end)/2, pivot)
	sIndex := start - 1
	eIndex := end + 1

	for {
		for {
			sIndex++
			fmt.Printf("Lower index increase %d -> %d \n", sIndex-1, sIndex)
			if sli[sIndex] >= pivot {
				fmt.Printf("Break lower index value %d -> Pivot value %d \n", sli[sIndex], pivot)
				break
			}
		}

		for {
			eIndex--
			fmt.Printf("Higher index decrease %d -> %d \n", eIndex+1, eIndex)
			if sli[eIndex] <= pivot {
				fmt.Printf("Break higher index value %d -> Pivot value %d \n", sli[eIndex], pivot)
				break
			}
		}

		if sIndex >= eIndex {
			return eIndex
		}

		sli[sIndex], sli[eIndex] = sli[eIndex], sli[sIndex]
	}
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

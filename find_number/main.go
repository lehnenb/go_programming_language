package main

import (
	"fmt"
	"os"
	"strconv"
)

func findIndex(list []int64, num int64) (int, bool) {
	for index, value := range list {
		if value == num {
			return index, true
		}
	}

	return 0, false
}

func main() {
	args := os.Args[1:]

	num, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		index, ok := findIndex([]int64{1, 2, 3, 4, 5, 6, 7, 8}, num)

		if ok {
			fmt.Printf("Value found at index: %d", index)
		} else {
			fmt.Print("Value not found")
		}
	}

	fmt.Print("Enter a valid number")
}

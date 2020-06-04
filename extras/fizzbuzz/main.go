package main

import (
	"fmt"
	"strconv"
)

const limit = 20

func getNumbers() []int {
	numbers := make([]int, limit)

	for i := 0; i < limit; i++ {
		numbers[i] = i + 1
	}

	return numbers
}

func getFizzBuzz(number int) string {
	divByThree := number%3 == 0
	divByFive := number%5 == 0

	if divByThree && divByFive {
		return "FizzBuzz"
	}

	if divByThree {
		return "Fizz"
	}

	if divByFive {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func getFizzBuzzes(numbers []int) []string {
	var fizzBuzz []string

	for _, number := range numbers {
		fizzBuzz = append(fizzBuzz, getFizzBuzz(number))
	}

	return fizzBuzz
}

func main() {
	numbers := getNumbers()
	fizzBuzzes := getFizzBuzzes(numbers)

	for _, fb := range fizzBuzzes {
		fmt.Println(fb)
	}
}

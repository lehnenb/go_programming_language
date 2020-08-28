package sum

// Sum a bunch of numbas
func Sum(list []int) int {
	sum := 0

	for _, value := range list {
		sum += value
	}

	return sum
}

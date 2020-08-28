package sumrecursive

// SummedValues encapuslates number list and accumulated value into single type
type SummedValues = struct {
	numbers []int
	acc     int
}

// SumRecursive sums a bunch of numbas
func SumRecursive(nums SummedValues) int {
	nums.acc += nums.numbers[0]

	if len(nums.numbers) == 1 {
		return nums.acc
	}

	nums.numbers = nums.numbers[1:]
	return SumRecursive(nums)
}

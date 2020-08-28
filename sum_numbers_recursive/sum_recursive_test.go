package sumrecursive

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		numbers  SummedValues
		expected int
	}{
		{SummedValues{[]int{1}, 0}, 1},
		{SummedValues{[]int{1, 1}, 0}, 2},
	}

	for _, caseValues := range cases {
		t.Run(fmt.Sprintf("(%v)", caseValues.numbers), func(t *testing.T) {
			got := SumRecursive(caseValues.numbers)

			if got != caseValues.expected {
				t.Fatalf("Sum() = %v; expected %v", got, caseValues.expected)
			}
		})
	}
}

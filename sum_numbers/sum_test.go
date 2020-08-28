package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{2, 2}, 4},
	}

	for _, caseValues := range cases {
		t.Run(fmt.Sprintf("(%v)", caseValues.numbers), func(t *testing.T) {
			got := Sum(caseValues.numbers)

			if got != caseValues.expected {
				t.Fatalf("Sum() = %v; expected %v", got, caseValues.expected)
			}
		})
	}
}

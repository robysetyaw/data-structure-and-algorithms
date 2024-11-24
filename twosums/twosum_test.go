package twosums

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	runTestCase(t, []int{}, 9, nil, "should return nil if given array is empty")
	runTestCase(t, []int{1, 2}, 3, []int{0, 1}, "should return 0,1 when given 1,2 and target is 3")
	runTestCase(t, []int{2, 3, 1}, 3, []int{2, 0}, "should return 2,0 when given 2,3,1 and target is 3")
	runTestCase(t, []int{1, 2, 3}, 4, []int{0, 2}, "should return 0,2 when given 1,2,3 and target is 4")
	runTestCase(t, []int{4, 2, 3}, 6, []int{1, 0}, "should return 1,0 when given 4,2,3 and target is 6")
	runTestCase(t, []int{4, 2, 3, 1, 5}, 9, []int{0, 4}, "should return 1,0 when given 4,2,3,1,5 and target is 9")
	runTestCase(t, []int{5, 2, 3, 1, 4}, 9, []int{4, 0}, "should return 1,0 when given 4,2,3,1,5 and target is 9")
	runTestCase(t, []int{2, 2, 5, 1, 4}, 9, []int{4, 2}, "should return 1,0 when given 4,2,3,1,5 and target is 9")
}

func runTestCase(t *testing.T, nums []int, target int, want []int, description string) {
	t.Run(description, func(t *testing.T) {
		got := TwoSum(nums, target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Failed: %s. TwoSum(%v, %d) = %v; want %v", description, nums, target, got, want)
		}
	})
}

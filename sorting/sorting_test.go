package sorting

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	runTestCase(t, []int{8, 2, 3, 5, 4, 1}, []int{1, 2, 3, 4, 5, 8}, "should return 1,2,3,4,5,8 when given 8,2,3,5,4,1")
	runTestCase(t, []int{}, []int{}, "should return nil when given blank")
	runTestCase(t, []int{1}, []int{1}, "should return []int{1} when given []int{1}")
	runTestCase(t, []int{2, 1}, []int{1, 2}, "should return []int{1,2} when given []int{2,1}")
	runTestCase(t, []int{2, 1, 3}, []int{1, 2, 3}, "should return []int{2,1,3} when given []int{1,2,3}")
	runTestCase(t, []int{2, 1, 3, 7, 3}, []int{1, 2, 3, 3, 7}, "should return []int{2,1,3,7,3} when given []int{1,2,3,3,7}")
	runTestCase(t, []int{2, 10, 3, 7, 8, 10}, []int{2, 3, 7, 8, 10, 10}, "should return []int{2,1,3,7,3} when given []int{1,2,3,3,7}")
}

func runTestCase(t *testing.T, nums []int, want []int, description string) {
	t.Run(description, func(t *testing.T) {
		got := Sorting(nums)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s. Sorting(%v) = %v; want %v", description, nums, got, want)
		}
	})
}

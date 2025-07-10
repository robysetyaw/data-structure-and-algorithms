package array

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	runTestFindMedianSortedArrays(t, []int{0}, []int{0}, float64(0), "should return 0 when input [0], [0]")
	runTestFindMedianSortedArrays(t, []int{0}, []int{2}, float64(1), "should return 1 when input [0], [2]")
	runTestFindMedianSortedArrays(t, []int{0}, []int{2, 4}, float64(2), "should return 2 when input [0], [2,4]")
	runTestFindMedianSortedArrays(t, []int{1, 2}, []int{3, 4}, float64(2.5), "should return 2 when input [0], [2,4]")
	runTestFindMedianSortedArrays(t, []int{100000}, []int{100001}, float64(100000.5), "should return 100000.50000 when input [100000], [100001]")
	runTestFindMedianSortedArrays(t, []int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, float64(9), "should return 9 when input [1,2,3,4,5], [6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]")
}

func runTestFindMedianSortedArrays(t *testing.T, nums1 []int, nums2 []int, expected float64, description string) {
	t.Run(description, func(t *testing.T) {
		got := findMedianSortedArrays(nums1, nums2)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v, input %v,%v", expected, got, nums1, nums2)
		}
	})
}

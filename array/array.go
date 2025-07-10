package array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if 3 > (len(nums1) + len(nums2)) {
		if 1 > len(nums1) {
			if len(nums2) > 1 {
				return medSortArr(nums2)
			}
			return float64(nums2[0])
		} else if 1 > len(nums2) {
			if len(nums1) > 1 {
				return medSortArr(nums1)
			}
			return float64(nums1[0])
		}
		return (float64(nums1[0]) + float64(nums2[0])) / 2.00
	}
	newNums := append(nums1, nums2...)
	v := medSortArr(newNums)
	return v
}

func medSortArr(newNums []int) float64 {
	sortedNums := quickSort(newNums, 0, len(newNums)-1)
	midNum := len(sortedNums) / 2
	if len(sortedNums)%2 == 0 {
		n1 := float64(sortedNums[midNum])
		n2 := float64(sortedNums[midNum-1])
		return (n1 + n2) / 2.00
	} else {
		return float64(sortedNums[midNum])
	}
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

package sorting

const (
	QUICKSORT           string = "QUICKSORT"
	BUBBLESORT          string = "BUBBLESORT"
	BUBBLESORTFOR       string = "BUBBLESORTFOR"
	BUBBLESORTRECURSIVE string = "BUBBLESORTRECURSIVE"
)

func Sorting(nums []int) []int {
	// return bubbleSort(nums)
	algorithm := QUICKSORT

	switch algorithm {
	case QUICKSORT:
		return quickSort(nums)
	case BUBBLESORT:
		return bubbleSort(nums)
	default:
		return quickSort(nums)
	}
}

func quickSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	pivotIndex := len(nums) / 2
	pivot := nums[pivotIndex]

	var left []int
	var right []int

	for i, v := range nums {
		if i == pivotIndex {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	result := []int{}
	result = append(result, sortedLeft...)
	result = append(result, pivot)
	result = append(result, sortedRight...)

	return result
}

func bubbleSort(nums []int) []int {
	method := BUBBLESORTRECURSIVE
	switch method {
	case BUBBLESORTRECURSIVE:
		return bubbleSortRecursive(nums, 0, 1)
	case BUBBLESORTFOR:
		return bubbleSortFor(nums)
	default:
		return bubbleSortRecursive(nums, 0, 1)
	}
}

func bubbleSortFor(nums []int) []int {
	i := 0
	for i < (len(nums)) {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
	}

	return nums
}

func bubbleSortRecursive(nums []int, firstIndex int, secondIndex int) []int {

	if firstIndex >= len(nums) {
		return nums
	}

	if secondIndex >= len(nums) {
		return bubbleSortRecursive(nums, firstIndex+1, firstIndex+2)
	}

	if nums[firstIndex] > nums[secondIndex] {
		nums[firstIndex], nums[secondIndex] = nums[secondIndex], nums[firstIndex]
	}

	return bubbleSortRecursive(nums, firstIndex, secondIndex+1)
}

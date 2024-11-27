package sorting

func Sorting(nums []int) []int {
	return sortingRecursive(nums)
}

func sortingFor(nums []int, ) []int {
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
func sortingRecursive(nums []int) []int {
	sortingRecursivePrivate(nums, 0, 1)

	return nums

}

func sortingRecursivePrivate(nums []int, firstIndex int, secondIndex int) []int {
	
	if firstIndex >= len(nums) {
		return nums
	}

	if secondIndex >= len(nums) {
		return sortingRecursivePrivate(nums, firstIndex+1, firstIndex+2)
	}

	if nums[firstIndex] > nums[secondIndex] {
        nums[firstIndex], nums[secondIndex] = nums[secondIndex], nums[firstIndex]
    }

	return sortingRecursivePrivate(nums, firstIndex, secondIndex+1)
}

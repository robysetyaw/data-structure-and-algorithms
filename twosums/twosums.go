package twosums

func TwoSum(nums []int, target int) []int {
	returnValue := twoSumWithMap(nums, target)
	return returnValue
}

func twoSumWithMap(nums []int, target int) []int {
	hashNum := map[int]int{}

	for i, num := range nums {
		hashNum[num] = i
	}

	for i, num := range nums {
		secondNumber := target - nums[i]
		value, isExist := hashNum[secondNumber]
		if isExist && secondNumber > num && value != i {
			return []int{i, value}
		} else if isExist && secondNumber < num {
			return []int{value, i}
		}
	}

	return nil
}

func twoSumWithArray(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		firstNumber := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondNumber := nums[j]
			if (firstNumber+secondNumber == target) && (firstNumber < secondNumber) {
				return []int{i, j}
			} else if (firstNumber+secondNumber == target) && (firstNumber > secondNumber) {
				return []int{j, i}
			}
		}
	}
	return nil
}

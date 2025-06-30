package stack

func ValidParentheses(input string) bool {
	var parentheses = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack = Stack{}
	if input == "" {
		return true
	}
	for _, v := range input {
		if v == '(' || v == '{' || v == '[' {
			stack.Push(v)
		} else {
			if stack.IsEmpty() {
				return false
			}
			opener, _ := stack.Top()
			if opener == parentheses[v] {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func LongestParentheses(input string) int {
	var maxLen = 0
	var stack = IntStack{}
	stack.Push(-1)
	for i, v := range input {
		if v == '(' || v == '{' || v == '[' {
			stack.Push(i)
		} else if v == ')' || v == '}' || v == ']' {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				length := i - stack.Top()
				if length > maxLen {
					maxLen = length
				}
			}
		}

	}
	return maxLen
}

func TrappingRainWater(input []int) int {
	total := 0
	finalTotal := 0
	left := 0
	right := 0
	if 3 > len(input) {
		return total
	} else {
		for i, v := range input {
			left = 0
			right = 0
			if i == 0 || i == (len(input)-1) {
				continue
			} else {

				for leftMax := i; leftMax >= 0; leftMax-- {
					if left < input[leftMax] {
						left = input[leftMax]
					}
				}

				for rigtMax := i; rigtMax < len(input); rigtMax++ {
					if right < input[rigtMax] {
						right = input[rigtMax]
					}
				}

				minHeight := min(left, right)
				total = minHeight - v

				if total >= 0 {
					finalTotal += total
				}

			}
		}
		return finalTotal
	}
}

func TrappingRainWaterWithStack(input []int) int {
	total := 0
	left := 0
	right := 0
	var stack = IntStack{}
	if len(input) < 3 {
		return 0
	}

	for i := 0; i < len(input); i++ {
		// fmt.Print("index ", i)
		for !stack.IsEmpty() && input[i] > input[stack.Top()] {
			bottomIndex := stack.Pop()
			bottom := input[bottomIndex]
			if stack.IsEmpty() {
				break
			}
			left = input[stack.Top()]
			right = input[i]
			distance := i - stack.Top() - 1
			bound := min(left, right) - bottom
			total = distance * bound
			// fmt.Printf(" valid: left = %d, right = %d, bottom, = %d, distance, %d, bound = %d, total = %d", left, right, distance, bottom, bound, total)
		}

		// fmt.Print(" stack = ", stack)
		// fmt.Println("")
		stack.Push(i)

	}
	return total
}

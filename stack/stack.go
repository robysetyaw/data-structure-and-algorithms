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

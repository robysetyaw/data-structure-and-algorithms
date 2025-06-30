package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	runValidParentheses(t)
	runLongestParentheses(t)
	runTrappingRainWater(t)
}

func runTrappingRainWater(t *testing.T) {
	runTestingTrappingRainWater(t, []int{0}, 0, "should return 0 when input = [0]")
	runTestingTrappingRainWater(t, []int{0, 1}, 0, "should return 0 when input = [0, 1]")
	runTestingTrappingRainWater(t, []int{0, 1, 2}, 0, "should return 0 when input = [0, 1, 2]")
	runTestingTrappingRainWater(t, []int{1, 0, 2}, 1, "should return 1 when input = [1, 0, 2]")
	runTestingTrappingRainWater(t, []int{2, 0, 2}, 2, "should return 2 when input = [2, 0, 2]")
	runTestingTrappingRainWater(t, []int{2, 0, 1}, 1, "should return 1 when input = [2, 0, 1]")
	runTestingTrappingRainWater(t, []int{2, 0, 1, 0, 1}, 2, "should return 2 when input = [2, 0, 1, 0, 1]")
	runTestingTrappingRainWater(t, []int{3, 0, 1, 0, 2}, 5, "should return 5 when input = [3, 0, 1, 0, 2]")
	runTestingTrappingRainWater(t, []int{3, 0, 0, 2, 0, 3}, 10, "should return 6 when input = [3, 0, 0, 2, 0, 3]")
	runTestingTrappingRainWater(t, []int{4, 2, 0, 3, 2, 5}, 9, "should return 9 when input = [4, 2, 0, 3, 2, 5]")
	runTestingTrappingRainWater(t, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6, "should return 6 when input = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]")
}

func runLongestParentheses(t *testing.T) {
	runTestingLongestParentheses(t, "", 0, "should return 0 when input = ''")
	runTestingLongestParentheses(t, "()}", 2, "should return 2 when input = '()}'")
	runTestingLongestParentheses(t, "(){}", 4, "should return 4 when input = '(){}'")
	runTestingLongestParentheses(t, ")()())", 4, "should return 4 when input = ')()())'")
	runTestingLongestParentheses(t, "()(()", 2, "should return 2 when input = '()(()'")
	runTestingLongestParentheses(t, "()(()[][{}}", 8, "should return 2 when input = '()(()'")
}

func runValidParentheses(t *testing.T) {
	runTestingValidParentheses(t, "", true, "should return true when input = ''")
	runTestingValidParentheses(t, "()", true, "should return true when input = '()'")
	runTestingValidParentheses(t, "[", false, "should return false when input = '['")
	runTestingValidParentheses(t, "(}", false, "should return false when input = '(}'")
	runTestingValidParentheses(t, "(){}[]", true, "should return true when input = '(){}[]'")
	runTestingValidParentheses(t, "({[]})", true, "should return true when input = '({[]})'")
	runTestingValidParentheses(t, "({]})", false, "should return false when input = '({]})'")
	runTestingValidParentheses(t, "({})({})", true, "should return true when input = '({})({})'")
}

func runTestingValidParentheses(t *testing.T, input string, output bool, description string) {
	t.Run(description, func(t *testing.T) {
		got := ValidParentheses(input)
		if !reflect.DeepEqual(got, output) {
			t.Errorf("Failed : %s, ValidParentheses(%v) = %v; want %v", description, input, got, output)
		}
	})
}

func runTestingLongestParentheses(t *testing.T, input string, output int, description string) {
	t.Run(description, func(t *testing.T) {
		got := LongestParentheses(input)
		if !reflect.DeepEqual(got, output) {
			t.Errorf("Failed : %s, LongestParentheses(%v) = %v; want %v", description, input, got, output)
		}
	})
}

func runTestingTrappingRainWater(t *testing.T, input []int, output int, description string) {
	t.Run(description, func(t *testing.T) {
		got := TrappingRainWaterWithStack(input)
		if !reflect.DeepEqual(got, output) {
			t.Errorf("Failed : %s, TrappingRainWater(%v) = %v; want %v", description, input, got, output)
		}
	})
}

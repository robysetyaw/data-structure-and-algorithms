package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	runValidParentheses(t)
	runLongestParentheses(t)
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

package string

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	runTestLengthOfLongestSubstring(t, "should return 0 when input = ''", "", 0)
	runTestLengthOfLongestSubstring(t, "should return 1 when input = 'a'", "a", 1)
	runTestLengthOfLongestSubstring(t, "should return 1 when input = 'aaaa'", "aaaa", 1)
	runTestLengthOfLongestSubstring(t, "should return 3 when input = 'aaaabc'", "aaaabc", 3)
}

func runTestLengthOfLongestSubstring(t *testing.T, description string, input string, expected int) {
	t.Run(description, func(t *testing.T) {
		got := LengthOfLongestSubstring(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("%s, got -> lengthOfLongestSubstring(%s) = %v, want = %v ", description, input, got, expected)
		}
	})
}

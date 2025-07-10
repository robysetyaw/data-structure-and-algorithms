package string

import (
	"fmt"
)

func LengthOfLongestSubstring(input string) int {
	// bruteForce(input)
	// slidingWindow(input)
	return slidingWindow(input)
}

func bruteForce(input string) int {
	fmt.Println("Brute force start ..")
	maxLength := 0
	n := len(input)
	fmt.Println("maxlength = ", maxLength)
	fmt.Println("n = ", n)

	for i := 0; i < n; i++ {
		fmt.Println("1st for")
		fmt.Println("i = ", i)
		seen := make(map[rune]bool)
		fmt.Println("seen = ", seen)
		currentLength := 0
		fmt.Println("currentLength = ", currentLength)

		for j := i; j < n; j++ {
			fmt.Println("2nd for")
			fmt.Println("j = ", j)
			ch := rune(input[j])
			fmt.Println("char = ", ch)
			if seen[ch] {
				fmt.Println("CH SEENED, BREAK 2nd for")
				break
			}
			fmt.Println("CH NOT SEENED, set seen[ch]==true")
			seen[ch] = true
			currentLength++
		}

		fmt.Println("Check currentLength > maxlength")
		if currentLength > maxLength {
			fmt.Println("true condition")
			maxLength = currentLength
		}
	}
	fmt.Println("Ended bruteforce")
	return maxLength
}

func slidingWindow(input string) int {
	start := 0
	maxLength := 0
	lastSeen := map[rune]int{}

	fmt.Println("sliding window start ..")
	fmt.Println("start =", start)
	fmt.Println("maxLength =", maxLength)
	fmt.Println("map lastSeen =", lastSeen)

	for i, ch := range input {
		fmt.Println("i = ", i)
		fmt.Println("ch = ", ch)
		idx, found := lastSeen[ch]
		fmt.Println("idx = ", idx)
		fmt.Println("found = ", found)

		fmt.Println("Check found & idx >= start")
		if found && idx >= start {
			fmt.Println("true condition")
			start = idx + 1
		}
		currentLength := i - start + 1
		fmt.Println("currentLength := i - start + 1 = ", currentLength)
		fmt.Println("check currentLength > maxlength")
		if currentLength > maxLength {
			fmt.Println("true condition")
			maxLength = currentLength
		}
		lastSeen[ch] = i
		fmt.Println("lastSeen[ch] = i", lastSeen[ch], " + ", i)
	}
	fmt.Println("Ended sliding window, maxlength = ", maxLength)
	return maxLength
}

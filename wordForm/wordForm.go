package wordform

func formWord(kata1 string, kata2 string) bool {
	// saya membuat maps dari character yang ada 
	charCount := make(map[rune]int)
	
	for _, char := range kata1 {
		charCount[char]++
	}

	for _, char := range kata2 {
		if charCount[char] == 0 {
			return false
		}
		charCount[char]--
	}

	return true
}
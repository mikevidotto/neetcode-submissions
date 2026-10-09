func isAnagram(s string, t string) bool {
	freq_map := make(map[rune]int)

	for _, char := range s {
		freq_map[char]++
	}

	for _, char := range t {
		freq_map[char]--
	}

	for _, freq := range freq_map {
		if freq != 0 {
			return false
		}
	}

	return true
}

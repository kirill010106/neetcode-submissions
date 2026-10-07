func isAnagram(s string, t string) bool {
	letters := [26]int{}
	for i := range s {
		letters[s[i]-97]++
	}
	for i := range t {
		letters[t[i]-97]--
	}
	for i := range letters {
		if letters[i] != 0 {
			return false
		}
	}
	return true
}

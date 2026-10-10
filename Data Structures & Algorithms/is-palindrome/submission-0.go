func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for !(unicode.IsLetter(rune(s[left])) || unicode.IsDigit(rune(s[left]))) && left < right {
			left++
		}
		for !(unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right]))) && left < right {
			right--
		}
		if strings.ToLower(string(s[left])) == strings.ToLower(string(s[right])) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

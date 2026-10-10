func isValid(s string) bool {
    stack := []byte{}
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if openBracket, ok := m[ch]; ok { // закрывающая
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) - 1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// открывающая
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

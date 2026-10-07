import "slices"

func isAnagram(s string, t string) bool {
	r1, r2 := []rune(s), []rune(t)
	slices.Sort(r1)
	slices.Sort(r2)
	s1, s2 := string(r1), string(r2)
	return s1 == s2
}

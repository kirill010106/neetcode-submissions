import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	r1, r2 := []rune(s), []rune(t)
	slices.Sort(r1)
	slices.Sort(r2)
	return slices.Equal(r1, r2)
}

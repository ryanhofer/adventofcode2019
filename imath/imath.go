package imath

// Cmp returns the ordering of a, b as -1, 0, +1.
func Cmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

package imath

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

// GCD returns the greatest common divisor of a, b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a, b.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

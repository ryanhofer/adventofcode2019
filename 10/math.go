package main

import "github.com/ryanhofer/adventofcode2019/imath"

func norm(x, y int) (int, int) {
	if x == 0 && y == 0 {
		return x, y
	}
	if x == 0 {
		return x, y / imath.Abs(y)
	}
	if y == 0 {
		return x / imath.Abs(x), y
	}
	d := gcd(imath.Abs(x), imath.Abs(y))
	return x / d, y / d
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

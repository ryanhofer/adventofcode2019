package main

func norm(x, y int) (int, int) {
	if x == 0 && y == 0 {
		return x, y
	}
	if x == 0 {
		return x, y / abs(y)
	}
	if y == 0 {
		return x / abs(x), y
	}
	d := gcd(abs(x), abs(y))
	return x / d, y / d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

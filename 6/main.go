package main

import (
	"fmt"
	"strings"

	"github.com/ryanhofer/adventofcode2019/input"
)

var orbits map[string]string

func main() {
	orbits = make(map[string]string)

	for line := range input.Lines() {
		parts := strings.Split(line, ")")
		parent := parts[0]
		child := parts[1]
		orbits[child] = parent
	}

	numOrbits := 0
	for child := range orbits {
		numOrbits += countOrbits(child)
	}
	fmt.Println("Part 1:", numOrbits)

	a := orbits["YOU"]
	b := orbits["SAN"]
	fmt.Println("Part 2:", distance(a, b))
}

func countOrbits(child string) int {
	parent, ok := orbits[child]
	if !ok {
		return 0
	}
	return 1 + countOrbits(parent)
}

func distance(start, end string) int {
	a := ancestors(start)
	b := ancestors(end)

	numCommonAncestors := 0
	for i := 0; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
		numCommonAncestors++
	}

	commonDistance := numCommonAncestors - 1
	return len(a) + len(b) - 2*commonDistance
}

func ancestors(child string) []string {
	result := []string{}
	parent, ok := orbits[child]
	for ok {
		result = append(result, parent)
		parent, ok = orbits[parent]
	}
	reverse(result)
	return result
}

func reverse(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

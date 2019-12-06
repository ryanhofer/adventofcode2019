package main

import (
	"fmt"
	"strconv"

	"github.com/ryanhofer/adventofcode2019/input"
)

func main() {
	var min int
	var max int
	fmt.Sscanf(input.Contents(), "%d-%d", &min, &max)

	numPasswords := 0
	numPasswordsStrict := 0
	for x := min; x <= max; x++ {
		if !decreasing(x) {
			continue
		}
		if doubles(x) {
			numPasswords++
		}
		if doublesStrict(x) {
			numPasswordsStrict++
		}
	}

	fmt.Println("Part 1:", numPasswords)
	fmt.Println("Part 2:", numPasswordsStrict)
}

func doubles(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i+1 < len(s); i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func doublesStrict(x int) bool {
	s := strconv.Itoa(x)

	r := 0
	runs := [][]byte{{s[0]}}

	for i := 1; i < len(s); i++ {
		last := runs[r][len(runs[r])-1]
		if s[i] == last {
			runs[r] = append(runs[r], s[i])
		} else {
			runs = append(runs, []byte{s[i]})
			r++
		}
	}

	for _, run := range runs {
		if len(run) == 2 {
			return true
		}
	}
	return false
}

func decreasing(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i+1 < len(s); i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

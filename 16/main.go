package main

import (
	"fmt"
	"strconv"

	"github.com/ryanhofer/adventofcode2019/imath"
	"github.com/ryanhofer/adventofcode2019/input"
)

var in, out []int
var pattern = []int{0, 1, 0, -1}

func main() {
	part1()
	part2()
}

func part1() {
	in, out = nil, nil
	for _, r := range input.Contents() {
		v, err := strconv.Atoi(string(r))
		check(err)
		in = append(in, v)
	}
	out = make([]int, len(in))

	for i := 0; i < 100; i++ {
		for r := range in {
			out[r] = 0
			for c := range in {
				out[r] += in[c] * p(r, c)
			}
			out[r] %= 10
			out[r] = imath.Abs(out[r])
		}
		in, out = out, in
	}

	s := ""
	for i := 0; i < 8; i++ {
		s += strconv.Itoa(in[i])
	}
	fmt.Println("Part 1:", s)
}

func p(r, c int) int {
	n := len(pattern)
	i := (c + 1) / (r + 1)
	return pattern[i%n]
}

func part2() {
	in, out = nil, nil
	var a []int
	contents := input.Contents()
	for _, r := range contents {
		v, err := strconv.Atoi(string(r))
		check(err)
		a = append(a, v)
	}
	for i := 0; i < 10000; i++ {
		in = append(in, a...)
	}
	out = make([]int, len(in))

	offset, err := strconv.Atoi(contents[:7])
	check(err)

	for i := 0; i < 100; i++ {
		sum := 0
		for r := len(in) - 1; r >= offset; r-- {
			sum += in[r]
			sum %= 10
			out[r] = sum
		}
		in, out = out, in
	}

	s := ""
	for i := 0; i < 8; i++ {
		s += strconv.Itoa(in[offset+i])
	}
	fmt.Println("Part 2:", s)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

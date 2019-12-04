package main

import "fmt"

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func str2dir(str string) Direction {
	switch str {
	case "U":
		return Up
	case "R":
		return Right
	case "D":
		return Down
	case "L":
		return Left
	default:
		panic(fmt.Sprintf("unexpected input %q", str))
	}
}

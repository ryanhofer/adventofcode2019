package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanhofer/adventofcode2019/input"
)

const width = 25
const height = 6

type Color uint8

const (
	Black       Color = 0
	White       Color = 1
	Transparent Color = 2
)

func (c Color) String() string {
	switch c {
	case Black:
		return " "
	case White:
		return "#"
	default:
		return "?"
	}
}

type Layer [height][width]Color

func main() {
	data := input.Contents()
	numLayers := len(data) / (width * height)
	layers := make([]Layer, numLayers)

	x, y, k := 0, 0, 0
	for _, s := range strings.Split(data, "") {
		value, err := strconv.Atoi(s)
		check(err)

		layers[k][y][x] = Color(value)

		x++
		if x >= width {
			x = 0
			y++
		}
		if y >= height {
			y = 0
			k++
		}
	}

	// find layer with fewest zeros
	bestCount := uint(width * height)
	bestK := 0
	for k := range layers {
		count := countPixels(0, &layers[k])
		if count < bestCount {
			bestCount = count
			bestK = k
		}
	}

	onesCount := countPixels(1, &layers[bestK])
	twosCount := countPixels(2, &layers[bestK])
	fmt.Println("Part 1:", onesCount*twosCount)

	fmt.Println("Part 2:")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			k := 0
			for layers[k][y][x] == Transparent {
				k++
			}
			fmt.Print(layers[k][y][x])
		}
		fmt.Println()
	}
}

func countPixels(match Color, layer *Layer) (count uint) {
	for y := range layer {
		for x := range layer[y] {
			if layer[y][x] == match {
				count++
			}
		}
	}
	return count
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

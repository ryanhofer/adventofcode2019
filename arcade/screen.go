package arcade

import "fmt"

type Screen map[Coord]Tile

func (s Screen) SetTile(x, y int, tile Tile) {
	c := Coord{x, y}
	s[c] = tile
}

func (s Screen) Bounds() (min, max Coord) {
	for c := range s {
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
	}
	return min, max
}

func (s Screen) Print() {
	min, max := s.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			c := Coord{x, y}
			fmt.Print(s[c])
		}
		fmt.Println()
	}
}

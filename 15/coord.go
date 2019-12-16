package main

type Coord struct {
	X, Y int
}

func (c Coord) Translate(dir Dir) Coord {
	switch dir {
	case North:
		c.Y--
	case South:
		c.Y++
	case West:
		c.X--
	case East:
		c.X++
	}
	return c
}

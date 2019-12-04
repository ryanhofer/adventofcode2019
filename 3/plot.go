package main

type Plot map[Point]int

func Draw(instructions []Instruction) Plot {
	x := 0
	y := 0
	totalDistance := 0

	plot := Plot{}
	plot.Mark(x, y, totalDistance)

	for _, instr := range instructions {
		for i := 0; i < instr.Distance; i++ {
			switch instr.Direction {
			case Up:
				y++
			case Right:
				x++
			case Down:
				y--
			case Left:
				x--
			}

			totalDistance++
			plot.Mark(x, y, totalDistance)
		}
	}

	return plot
}

func (p Plot) Mark(x, y, distance int) {
	point := Point{
		X: x,
		Y: y,
	}
	p[point] = distance
}

func Intersect(a, b Plot) Plot {
	intersect := Plot{}
	for p, d1 := range a {
		d2, ok := b[p]
		if ok {
			intersect[p] = d1 + d2
		}
	}
	return intersect
}

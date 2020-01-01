package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/geom"
	"github.com/ryanhofer/adventofcode2019/input"
)

func main() {
	m := maze{}

	keys := map[rune]geom.Vec{}

	var c, entrance geom.Vec
	for line := range input.Lines() {
		for _, r := range line {
			t := tile(r)
			switch true {
			case t.isKey():
				keys[r] = c
			case t == Entrance:
				entrance = c
			}
			m[c] = t
			c.X++
		}
		c.X = 0
		c.Y++
	}

	fmt.Printf("entrance=%s\n", entrance)
	fmt.Printf("keys=%v\n", keys)

	// var wg sync.WaitGroup

	p := &pathfinder{
		m:   m,
		pos: entrance,
	}
	path, ok := p.solve(keys)
	if !ok {
		fmt.Println("NO PATH")
		return
	}
	fmt.Println("BEST:", path)
	fmt.Println("Part 2:", len(path))
}

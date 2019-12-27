package main

import "github.com/ryanhofer/adventofcode2019/imath"

func totalEnergy() int {
	totalEnergy := 0
	for i := range moons {
		totalEnergy += moons[i].energy()
	}
	return totalEnergy
}

func runSimulation(ticks int) {
	for t := 0; t < ticks; t++ {
		step()
	}
}

func findCycleLength() int {
	initial := []Moon{}
	initial = append(initial, moons...)

	type searchstate struct {
		found bool
		t     int
	}

	var cycle vec3

	t := 0
	for {
		step()
		t++

		var diff vec3
		for i := range moons {
			dp := initial[i].pos.cmp(moons[i].pos).negate()
			dv := initial[i].vel.cmp(moons[i].vel).negate()

			if dp.x != 0 || dv.x != 0 {
				diff.x = 1
			}
			if dp.y != 0 || dv.y != 0 {
				diff.y = 1
			}
			if dp.z != 0 || dv.z != 0 {
				diff.z = 1
			}
		}

		if diff.x == 0 && cycle.x == 0 {
			cycle.x = t
		}
		if diff.y == 0 && cycle.y == 0 {
			cycle.y = t
		}
		if diff.z == 0 && cycle.z == 0 {
			cycle.z = t
		}

		if cycle.x > 0 && cycle.y > 0 && cycle.z > 0 {
			break
		}
	}

	return imath.LCM(cycle.x, imath.LCM(cycle.y, cycle.z))
}

func step() {
	// gravity
	for i := range moons {
		for j := range moons {
			if i == j {
				continue
			}
			m1, m2 := &moons[i], &moons[j]
			dv := m1.pos.cmp(m2.pos).negate()
			m1.vel = m1.vel.add(dv)
		}
	}

	// velocity
	for i := range moons {
		m := &moons[i]
		m.pos = m.pos.add(m.vel)
	}
}

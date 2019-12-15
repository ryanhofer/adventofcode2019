package main

import (
	"fmt"
)

type Chemical string

const (
	Fuel Chemical = "FUEL"
	Ore  Chemical = "ORE"
)

type Reaction struct {
	Inputs map[Chemical]int
	Output Chemical
	Yield  int
}

var reactions map[Chemical]Reaction
var leftovers map[Chemical]int

func main() {
	loadReactionList()

	leftovers = map[Chemical]int{}
	oreCostPerFuel := produce(Fuel, 1)
	fmt.Println("Part 1:", oreCostPerFuel)

	leftovers = map[Chemical]int{}
	oreRemaining := 1_000_000_000_000
	totalFuelProduced := 0

	batchSize := oreRemaining / oreCostPerFuel
	for batchSize > 0 {
		// Copy leftovers so that we can roll back
		tmp := map[Chemical]int{}
		for k, v := range leftovers {
			if v > 0 {
				tmp[k] = v
			}
		}

		// Attempt to produce batch
		oreCost := produce(Fuel, batchSize)
		if oreCost > oreRemaining {
			// Failed; batchSize is too large
			leftovers = tmp
			batchSize /= 10
			continue
		}
		oreRemaining -= oreCost
		totalFuelProduced += batchSize
	}

	fmt.Println("Part 2:", totalFuelProduced)
}

func produce(want Chemical, quantity int) int {
	if want == Ore {
		return quantity
	}

	if leftovers[want] >= quantity {
		leftovers[want] -= quantity
		return 0
	}

	if leftovers[want] > 0 {
		quantity -= leftovers[want]
		leftovers[want] = 0
	}

	r := reactions[want]
	numReactions := (quantity + r.Yield - 1) / r.Yield

	oreCost := 0
	for chem, qty := range r.Inputs {
		oreCost += produce(chem, qty*numReactions)
	}

	leftovers[want] += r.Yield*numReactions - quantity
	return oreCost
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

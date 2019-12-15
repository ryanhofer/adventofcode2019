package main

import (
	"fmt"
	"strings"

	"github.com/ryanhofer/adventofcode2019/input"
)

func loadReactionList() {
	reactions = map[Chemical]Reaction{}

	for line := range input.Lines() {
		parts := strings.Split(line, " => ")

		chem, qty, err := parseChemQty(parts[1])
		check(err)

		reaction := Reaction{
			Inputs: map[Chemical]int{},
			Output: chem,
			Yield:  qty,
		}

		for _, inputStr := range strings.Split(parts[0], ", ") {
			chem, qty, err := parseChemQty(inputStr)
			check(err)
			reaction.Inputs[chem] = qty
		}

		reactions[reaction.Output] = reaction
	}
}

func parseChemQty(in string) (chem Chemical, qty int, err error) {
	_, err = fmt.Sscanf(in, "%d %s", &qty, &chem)
	return chem, qty, err
}

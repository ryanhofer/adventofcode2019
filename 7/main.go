package main

import (
	"fmt"
	"sync"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

const numAmplifiers = 5

func main() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	var bestPhase []int
	var bestSignal intcode.Word

	for _, phase := range permutations([]int{0, 1, 2, 3, 4}) {
		var wg sync.WaitGroup
		wg.Add(numAmplifiers)

		firstInput := make(chan intcode.Word)
		signalChan := firstInput

		for i := 0; i < numAmplifiers; i++ {
			phaseChan := make(chan intcode.Word)
			in := merge(phaseChan, signalChan)
			out := make(chan intcode.Word)

			go func(in <-chan intcode.Word, out chan<- intcode.Word) {
				cfg := &intcode.Config{
					Input:  in,
					Output: out,
				}
				_, err := intcode.Exec(program, cfg)
				check(err)
				wg.Done()
			}(in, out)

			phaseChan <- intcode.Word(phase[i])
			close(phaseChan)

			signalChan = out
		}

		firstInput <- 0

		// consume all output
		wg.Add(1)
		var signal intcode.Word
		go func() {
			for signal = range signalChan {
			}
			wg.Done()
		}()

		wg.Wait()

		// fmt.Println(phase, "->", signal)
		if signal > bestSignal {
			bestPhase = phase
			bestSignal = signal
		}
	}

	fmt.Println("BEST:", bestPhase, "->", bestSignal)
	fmt.Println("Part 1:", bestSignal)

	bestSignal = 0
	bestPhase = nil

	for _, phase := range permutations([]int{5, 6, 7, 8, 9}) {
		var wg sync.WaitGroup
		wg.Add(numAmplifiers)

		firstInput := make(chan intcode.Word)
		lastOutput := firstInput

		for i := 0; i < numAmplifiers; i++ {
			phaseChan := make(chan intcode.Word)
			in := merge(phaseChan, lastOutput)
			out := make(chan intcode.Word)

			go func(in <-chan intcode.Word, out chan<- intcode.Word) {
				cfg := &intcode.Config{
					Input:  in,
					Output: out,
				}
				_, err := intcode.Exec(program, cfg)
				check(err)
				wg.Done()
			}(in, out)

			phaseChan <- intcode.Word(phase[i])
			close(phaseChan)

			lastOutput = out
		}

		firstInput <- 0

		// consume all output
		wg.Add(1)
		var signal intcode.Word
		go func() {
			for signal = range lastOutput {
				// feedback loop
				firstInput <- signal
			}
			wg.Done()
		}()

		wg.Wait()

		// fmt.Println(phase, "->", signal)
		if signal > bestSignal {
			bestSignal = signal
			bestPhase = phase
		}
	}

	fmt.Println("BEST:", bestPhase, "->", bestSignal)
	fmt.Println("Part 2:", bestSignal)
}

func merge(chs ...<-chan intcode.Word) <-chan intcode.Word {
	merged := make(chan intcode.Word)
	for _, ch := range chs {
		go func(ch <-chan intcode.Word) {
			for v := range ch {
				merged <- v
			}
		}(ch)
	}
	return merged
}

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permutations(a []int) (result [][]int) {
	var generate func(int, []int)
	generate = func(k int, a []int) {
		if k == 1 {
			b := make([]int, len(a))
			copy(b, a)
			result = append(result, b)
		} else {
			generate(k-1, a)
			for i := 0; i < k-1; i++ {
				if k%2 == 0 {
					a[i], a[k-1] = a[k-1], a[i]
				} else {
					a[0], a[k-1] = a[k-1], a[0]
				}
				generate(k-1, a)
			}
		}
	}
	generate(len(a), a)
	return
}

func extract(phase, i int) int {
	const modulo = 5
	divisor := 1
	for i > 0 {
		divisor *= modulo
		i--
	}
	return (phase / divisor) % modulo
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

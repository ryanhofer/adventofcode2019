package intcode

func Spawn(program Program) (in chan<- Word, out <-chan Word, halt <-chan error) {
	inChan := make(chan Word)
	outChan := make(chan Word)
	haltChan := make(chan error)

	go func(in <-chan Word, out chan<- Word) {
		cfg := &Config{
			Input:  in,
			Output: out,
		}
		_, err := Exec(program, cfg)
		haltChan <- err
	}(inChan, outChan)

	return inChan, outChan, haltChan
}

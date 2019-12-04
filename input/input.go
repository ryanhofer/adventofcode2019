package input

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func File() *os.File {
	f, err := os.Open("input.txt")
	check(err)
	return f
}

func Contents() string {
	f := File()
	defer f.Close()

	input, err := ioutil.ReadAll(f)
	check(err)

	contents := string(input)
	contents = strings.ReplaceAll(contents, "\r\n", "\n")
	contents = strings.Trim(contents, "\n")

	return contents
}

func Lines() <-chan string {
	lines := make(chan string)
	f := File()

	go func() {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			lines <- scanner.Text()
		}

		close(lines)
	}()

	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

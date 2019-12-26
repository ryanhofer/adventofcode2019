package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"

	. "github.com/ryanhofer/adventofcode2019/geom"
)

var maze Maze
var robot Robot

func main() {
	part1()
	part2()
}

func part1() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	_, outChan, haltChan := intcode.Spawn(program)
	parseCameraFeed(outChan)
	printMaze()

	sum := 0
	for v := range maze {
		if maze.IsIntersection(v) {
			sum += v.X * v.Y
		}
	}
	fmt.Println("Part 1:", sum)

	<-haltChan
}

func part2() {
	path := findPath()

	m, a, b, c := compressPath(path)
	fmt.Println("Main:", m)
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("C:", c)

	program, err := intcode.Parse(input.Contents())
	check(err)

	// Wake up!
	program[0] = intcode.Word(2)

	inChan, outChan, haltChan := intcode.Spawn(program)

	cameraChan := make(chan intcode.Word)
	resultChan := make(chan intcode.Word)
	go func() {
		defer close(resultChan)
		defer close(cameraChan)
		for r := range outChan {
			if r < 0x100 {
				cameraChan <- r
			} else {
				resultChan <- r
				return
			}
		}
	}()

	go func() {
		parseCameraFeed(cameraChan)
	}()

	// Send movement routines to robot
	code := fmt.Sprintf("%s\n%s\n%s\n%s\n", m, a, b, c)
	for _, r := range code {
		inChan <- intcode.Word(r)
	}

	// Disable continuous video feed
	inChan <- 'n'
	inChan <- '\n'

	dustCollected := <-resultChan
	fmt.Println("Part 2:", dustCollected)

	<-haltChan
}

func findPath() []string {
	program, err := intcode.Parse(input.Contents())
	check(err)

	_, outChan, _ := intcode.Spawn(program)
	parseCameraFeed(outChan)

	// Find the path
	moves := []string{}

navigate:
	for {
		switch Scaffold {
		case maze[robot.Left()]:
			robot.TurnLeft()
			moves = append(moves, "L")
		case maze[robot.Right()]:
			robot.TurnRight()
			moves = append(moves, "R")
		default:
			break navigate
		}

		steps := 0
		for maze[robot.Forward()] == Scaffold {
			robot.MoveForward()
			steps++
		}
		moves = append(moves, strconv.Itoa(steps))
	}

	return moves
}

func compressPath(path []string) (m, a, b, c string) {
	// Brute force approach:
	// Try all non-overlapping substrings a,b,c of path

	const maxRoutineLength int = 20

	re := regexp.MustCompile("^([ABC],)*[ABC]$")
	p := flatten(path)

	for k1 := 1; k1 < len(path); k1++ {
		for i1 := 0; i1+k1 < maxRoutineLength/2; i1++ {
			a = flatten(path[i1 : i1+k1])
			if len(a) > maxRoutineLength {
				continue
			}
			p1 := strings.ReplaceAll(p, a, "A")

			for k2 := 1; k2 < maxRoutineLength/2; k2++ {
				for i2 := i1 + k1 + 1; i2+k2 < len(path); i2++ {
					b = flatten(path[i2 : i2+k2])
					if len(b) > maxRoutineLength {
						continue
					}
					p2 := strings.ReplaceAll(p1, b, "B")

					for k3 := 1; k3 < maxRoutineLength/2; k3++ {
						for i3 := i2 + k2 + 1; i3+k3 < len(path); i3++ {
							c = flatten(path[i3 : i3+k3])
							if len(c) > maxRoutineLength {
								continue
							}
							p3 := strings.ReplaceAll(p2, c, "C")
							if len(p3) > maxRoutineLength {
								continue
							}
							if re.MatchString(p3) {
								m = p3
								return
							}
						}
					}
				}
			}
		}
	}
	panic("no solution found")
}

func flatten(s []string) string {
	return strings.Join(s, ",")
}

func parseCameraFeed(outChan <-chan intcode.Word) {
	maze = Maze{}
	robot = Robot{}

	c := Vec{}
	for out := range outChan {
		r := rune(out)

		switch r {
		case '\n':
			break
		case OpenSpace:
			maze[c] = OpenSpace
		default:
			maze[c] = Scaffold
		}

		if isRobot(r) {
			var dir Cardinal
			switch r {
			case RobotN:
				dir = North
			case RobotE:
				dir = East
			case RobotS:
				dir = South
			case RobotW:
				dir = West
			}
			robot = Robot{
				pos: c,
				dir: dir,
			}
		}

		c.X++
		if r == '\n' {
			c.X = 0
			c.Y++
		}
	}
}

func printMaze() {
	var min, max Vec
	for c := range maze {
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

	w := bufio.NewWriter(os.Stdout)
	var c Vec
	for c.Y = min.Y; c.Y <= max.Y; c.Y++ {
		for c.X = min.X; c.X <= max.X; c.X++ {
			r, ok := maze[c]
			if !ok {
				r = OpenSpace
			}
			if c == robot.pos {
				if r == Scaffold {
					w.WriteRune(robot.Rune())
				} else {
					w.WriteRune(RobotFalling)
				}
			} else {
				w.WriteRune(r)
			}
		}
		w.WriteRune('\n')
	}
	w.Flush()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

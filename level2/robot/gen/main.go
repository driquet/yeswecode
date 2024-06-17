package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type coord struct {
	row int
	col int
}

func (c coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.row, c.col)
}

func generate(input, output string) error {
	data, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	// Split into lines
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	height := len(lines)
	width := len(lines[0])

	// Ensure each line is the same length
	for idx, line := range lines {
		if len(line) != width {
			return fmt.Errorf("line %d with incorrect width, got %d, want %d", idx, len(line), width)
		}
	}

	// List the positions we want with a trail
	want := make(map[coord]bool)
	for row, line := range lines {
		for col, r := range line {
			if r != ' ' {
				want[coord{
					row: row,
					col: col,
				}] = true
			}
		}
	}

	var str strings.Builder

	// Random starting point
	pos := coord{
		row: rand.Intn(height),
		col: rand.Intn(width),
	}
	var trail bool

	have := make(map[coord]bool)

	directions := []rune{'<', '>', '^', 'v'}
	moves := map[rune]coord{
		'<': {row: 0, col: -1},
		'>': {row: 0, col: 1},
		'^': {row: -1, col: 0},
		'v': {row: 1, col: 0},
	}
	var idx int

	// Continue until we have left a trail on all the wanted cells
	for len(want) != len(have) {
		// From the current position, randomly select a direction and a distance
		direction := directions[idx%len(directions)]
		move := moves[direction]
		distance := rand.Intn(42)

		for i := 0; i < distance; i++ {
			// Make sure the robot is not out of bounds
			next := coord{
				row: pos.row + move.row,
				col: pos.col + move.col,
			}
			if next.row < 0 || next.row >= height || next.col < 0 || next.col >= width {
				break
			}
			// Determine if we need to toggle the trail mode
			if (trail && !want[pos]) || (!trail && want[pos]) {
				trail = !trail
				str.WriteRune('!')
			}

			// Leave a trail
			if trail {
				have[pos] = true
				if !want[pos] {
					return fmt.Errorf("trail left on a cell that was not supposed to have it")
				}
			}

			// Move the robot
			pos = next
			str.WriteRune(direction)
		}

		idx++
	}

	return os.WriteFile(output, []byte(str.String()), 0600)
}

func main() {
	var pattern, output string

	pflag.StringVar(&pattern, "pattern", "pattern.txt", "representation of the pattern to find")
	pflag.StringVar(&output, "output", "input.txt", "name of the generated input file")
	pflag.Parse()

	if err := generate(pattern, output); err != nil {
		fmt.Printf("unable to generate input: %s\n", err)
		os.Exit(1)
	}
}

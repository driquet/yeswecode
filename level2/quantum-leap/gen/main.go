package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type point struct {
	row int
	col int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.col, p.row)
}

type qubit struct {
	pos      point
	velocity point
}

func randValue(from, to int) int {
	n := rand.Intn(to - from + 1)
	return from + n
}

func generate(input, output string) error {
	data, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	var bits []*qubit

	rows := strings.Split(string(data), "\n")
	for row, line := range rows {
		for col, r := range line {
			if r == '#' {
				bits = append(bits, &qubit{
					pos: point{
						row: row,
						col: col,
					},
					velocity: point{
						row: randValue(-3, 3),
						col: randValue(-3, 3),
					},
				})
			}
		}
	}

	// Simulate (backward)
	steps := randValue(30, 42)
	for i := 0; i < steps; i++ {
		for _, qb := range bits {
			qb.pos.row = qb.pos.row - qb.velocity.row
			qb.pos.col = qb.pos.col - qb.velocity.col
		}
	}

	// Write to file
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, qb := range bits {
		fmt.Fprintf(f, "%s %s\n", qb.pos, qb.velocity)
	}

	return nil
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

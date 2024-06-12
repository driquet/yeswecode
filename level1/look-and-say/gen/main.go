package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var count = 42

func generate(value int) string {
	curr := fmt.Sprintf("%d", value)
	for idx := 0; idx < count; idx++ {
		curr = nextTerm(curr)
	}

	return curr
}

func nextTerm(term string) string {
	if len(term) == 0 {
		return ""
	}

	var str strings.Builder

	curr := term[0]
	count := 1

	for idx := 1; idx < len(term); idx++ {
		if term[idx] == curr {
			count++
			continue
		}

		str.WriteString(fmt.Sprintf("%d%c", count, curr))

		curr = term[idx]
		count = 1
	}

	str.WriteString(fmt.Sprintf("%d%c", count, curr))

	return str.String()
}

func main() {
	var value int

	flag.IntVar(&value, "f", 1337, "flag value")
	flag.Parse()

	var output *os.File
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			fmt.Printf("error opening file %q: %v\n", filename, err)
			os.Exit(1)
		}
		defer f.Close()

		output = f
	} else {
		output = os.Stdout
	}

	input := generate(value)
	fmt.Fprint(output, input)
}

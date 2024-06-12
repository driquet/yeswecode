package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func generate(min int) (string, int) {
	var input strings.Builder

	var countParentheses, countCurly, countSquare int

	posParentheses := -1
	posCurly := -1
	posSquare := -1

	trailing := rand.Intn(min)

	for posParentheses == -1 || posCurly == -1 || posSquare == -1 || trailing > 0 {
		pos := input.Len()

		// Choose the next character
		choices := []rune{'(', '{', '['}

		// Closing characters can be picked only if they do not create unbalancedness before min
		for r, count := range map[rune]int{
			')': countParentheses,
			'}': countCurly,
			']': countSquare,
		} {
			if count == 0 && pos < min {
				break
			}
			choices = append(choices, r)
		}

		choice := rand.Intn(len(choices))
		r := choices[choice]

		switch r {
		case '(':
			countParentheses++
		case ')':
			countParentheses--
			if countParentheses < 0 && posParentheses == -1 {
				posParentheses = pos
			}
		case '{':
			countCurly++
		case '}':
			countCurly--
			if countCurly < 0 && posCurly == -1 {
				posCurly = pos
			}
		case '[':
			countSquare++
		case ']':
			countSquare--
			if countSquare < 0 && posSquare == -1 {
				posSquare = pos
			}
		}

		input.WriteRune(r)

		if posParentheses != -1 && posCurly != -1 && posSquare != -1 {
			trailing--
		}
	}

	fmt.Printf("par(%d) cur(%d) squ(%d)\n", posParentheses, posCurly, posSquare)
	flag := posParentheses * posCurly * posSquare
	return input.String(), flag
}

func main() {
	var min int

	flag.IntVar(&min, "m", 100, "minimum size of the generated challenge input")
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

	input, flag := generate(min)

	fmt.Printf("flag: %d\n", flag)
	fmt.Fprint(output, input)
}

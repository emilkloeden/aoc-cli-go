package main

import (
	"fmt"
	"os"
)

var prod = false
var testInput = ``

func main() {
	var input string = testInput
	if prod {
		input = loadInput()
	}
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}
	if args[1] == "1" {
		part1(input)
	} else if args[1] == "2" {
		part2(input)
	} else {
		panic("Must supply either '1' or '2' as part argument.")
	}

}

func loadInput() string {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func part1(input string) {
	panic("Part 1 not implemented.")
}

func part2(input string) {
	panic("Part 2 not implemented.")
}

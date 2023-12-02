package main

import (
	"bufio"
	"log"
	"os"
)

func part1(data *bufio.Scanner) {
}

func part2(data *bufio.Scanner) {
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--example" {
		example, err := readInput(InputFileExample1)
		if err != nil {
			log.Fatal(err)
		}
		part1(example)

		example, err = readInput(InputFileExample2)
		if err != nil {
			log.Fatal(err)
		}
		part2(example)
		return
	}

	input, err := readInput(InputFileReal)
	if err != nil {
		log.Fatal(err)
	}
	part1(input)

	input, err = readInput(InputFileReal)
	if err != nil {
		log.Fatal(err)
	}

	part2(input)
}

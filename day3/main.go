package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func is_valid(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) && c != '.' {
			return true
		}
	}

	return false
}

func part1(data *bufio.Scanner) {
	text := []string{}
	for data.Scan() {
		text = append(text, data.Text())
	}

	total := 0
	for lineIndex := 0; lineIndex < len(text); lineIndex++ {
		for charIndex := 0; charIndex < len(text[lineIndex]); charIndex++ {
			if unicode.IsDigit(rune(text[lineIndex][charIndex])) {
				startIndex := charIndex
				endIndex := charIndex

				for endIndex < len(text[lineIndex]) && unicode.IsDigit(rune(text[lineIndex][endIndex])) {
					endIndex++
				}

				charIndex = endIndex

				searchStart := startIndex
				searchEnd := endIndex

				if searchStart == 0 {
					searchStart += 1
				}

				if searchEnd == len(text[lineIndex]) {
					searchEnd -= 1
				}

				box := ""

				// top
				if lineIndex != 0 {
					box += text[lineIndex-1][searchStart-1 : searchEnd+1]
				}

				// current
				box += text[lineIndex][searchStart-1 : searchEnd+1]

				// bottom
				if lineIndex != len(text)-1 {
					box += text[lineIndex+1][searchStart-1 : searchEnd+1]
				}

				if is_valid(box) {
					num, err := strconv.Atoi(text[lineIndex][startIndex:endIndex])
					if err != nil {
						log.Fatal(err)
					}

					total += num
				}
			}
		}
	}
	fmt.Printf("Total: %d\n", total)
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(data *bufio.Scanner) {
	total := 0

	for data.Scan() {
		firstDigit := 0
		lastDigit := 0

		for _, c := range data.Text() {
			if digit, err := strconv.ParseInt(string(c), 10, 64); err == nil {
				if firstDigit == 0 {
					firstDigit = int(digit)
				}

				lastDigit = int(digit)
			}
		}

		total += lastDigit + (firstDigit * 10)
	}

	fmt.Printf("Part1 total: %d\n", total)
}

func part2(data *bufio.Scanner) {
	digits := map[string]int{
		"one":       1,
		"two":       2,
		"three":     3,
		"four":      4,
		"five":      5,
		"six":       6,
		"seven":     7,
		"eight":     8,
		"nine":      9,
		"twone":     21,
		"eightwo":   81,
		"nineight":  98,
		"eighthree": 83,
	}

	total := 0

	parsedText := []string{}

	for data.Scan() {
		line := data.Text()

		foundNumbers := []int{}

		for k, v := range digits {
			re := regexp.MustCompile(k)
			for _, _ = range re.FindAllIndex([]byte(line), -1) {
				foundNumbers = append(foundNumbers, v)
			}
		}

		parsedText = append(parsedText, line)
	}

	for _, line := range parsedText {
		firstDigit := 0
		lastDigit := 0

		for _, c := range line {
			digit, err := strconv.ParseInt(string(c), 10, 64)
			if err != nil {
				continue
			}

			if firstDigit == 0 {
				firstDigit = int(digit)
			}

			lastDigit = int(digit)
		}

		total += lastDigit + (firstDigit * 10)
	}

	fmt.Printf("Part2 total: %d\n", total)
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

	input, err := readInput(InputFileReal1)
	if err != nil {
		log.Fatal(err)
	}
	part1(input)

	input, err = readInput(InputFileReal2)
	if err != nil {
		log.Fatal(err)
	}
	part2(input)
}

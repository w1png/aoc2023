package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(data *bufio.Scanner) {
	total := 0

	for data.Scan() {
		card_number_draws := strings.Split(data.Text(), ":")
		// card_number, err := strconv.Atoi(strings.Split(card_number_draws[0], " ")[1])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		draws := strings.Split(strings.TrimSpace(card_number_draws[1]), "|")

		winning_cards := map[int]bool{}
		for _, draw := range strings.Split(strings.TrimSpace(draws[0]), " ") {
			if draw == "" {
				continue
			}
			draw_int, err := strconv.Atoi(draw)
			if err != nil {
				log.Fatal(err)
			}
			winning_cards[draw_int] = true
		}

		matches := 0

		for _, draw := range strings.Split(strings.TrimSpace(draws[1]), " ") {
			if draw == "" {
				continue
			}
			draw_int, err := strconv.Atoi(draw)
			if err != nil {
				log.Fatal(err)
			}

			if winning_cards[draw_int] {
				matches++
			}
		}

		total += int(math.Pow(2, float64(matches-1)))
	}

	fmt.Printf("Part1 total: %d\n", total)
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

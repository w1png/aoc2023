package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(data *bufio.Scanner) {
	max_red := 12
	max_green := 13
	max_blue := 14

	// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	total := 0

	for data.Scan() {
		line := data.Text()
		is_valid := true

		s := strings.Split(line, ":")
		gameIdString := s[0]
		gameDrawsString := strings.TrimSpace(s[1])

		id, err := strconv.Atoi(strings.Split(gameIdString, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		for _, draw := range strings.Split(gameDrawsString, ";") {
			draw = strings.TrimSpace(draw)
			var red, green, blue int
			for _, stone := range strings.Split(draw, ",") {
				stone = strings.TrimSpace(stone)
				d := strings.Split(stone, " ")
				amount, err := strconv.Atoi(strings.TrimSpace(d[0]))
				if err != nil {
					log.Fatal(err)
				}

				color := strings.TrimSpace(d[1])
				switch color {
				case "red":
					red = amount
				case "green":
					green = amount
				case "blue":
					blue = amount
				}
			}

			if red > max_red || green > max_green || blue > max_blue {
				is_valid = false
			}
		}

		if is_valid {
			total += id
		}
	}

	fmt.Printf("Part 1 result: %d\n", total)
}

func part2(data *bufio.Scanner) {
	total := 0

	for data.Scan() {
		line := data.Text()
		max_blue := 0
		max_red := 0
		max_green := 0

		s := strings.Split(line, ":")
		gameDrawsString := strings.TrimSpace(s[1])

		for _, draw := range strings.Split(gameDrawsString, ";") {
			draw = strings.TrimSpace(draw)
			var red, green, blue int
			for _, stone := range strings.Split(draw, ",") {
				stone = strings.TrimSpace(stone)
				d := strings.Split(stone, " ")
				amount, err := strconv.Atoi(strings.TrimSpace(d[0]))
				if err != nil {
					log.Fatal(err)
				}

				color := strings.TrimSpace(d[1])
				switch color {
				case "red":
					red = amount
				case "green":
					green = amount
				case "blue":
					blue = amount
				}
			}

			if max_red < red {
				max_red = red
			}

			if max_green < green {
				max_green = green
			}

			if max_blue < blue {
				max_blue = blue
			}
		}

		total += max_red * max_green * max_blue
	}

	fmt.Printf("Part 2 result: %d\n", total)
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

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
	parseValues := func(data string) ([]int, error) {
		output := []int{}
		for _, time := range strings.Split(strings.TrimSpace(strings.Split(data, ":")[1]), " ") {
			if time == "" {
				continue
			}

			timeInt, err := strconv.Atoi(time)
			if err != nil {
				return nil, err
			}

			output = append(output, timeInt)
		}

		return output, nil
	}

	total := 1
	var times, records []int
	var err error

	data.Scan()
	if times, err = parseValues(data.Text()); err != nil {
		log.Fatal(err)
	}
	data.Scan()
	if records, err = parseValues(data.Text()); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(records); i++ {
		record := records[i]
		time := times[i]

		totalMargin := 0
		for holdTime := 0; holdTime < time; holdTime++ {
			speed := holdTime

			if (time-holdTime)*speed > record {
				totalMargin++
			}
		}

		total *= totalMargin
	}

	fmt.Printf("Part1 total: %d\n", total)
}

func part2(data *bufio.Scanner) {
	parseValues := func(data string) string {
		output := ""

		for _, val := range strings.Split(strings.TrimSpace(strings.Split(data, ":")[1]), " ") {
			if val == "" {
				continue
			}

			output += val
		}
		return output
	}

	total := 1
	var time, record int
	var err error

	data.Scan()
	if time, err = strconv.Atoi(parseValues(data.Text())); err != nil {
		log.Fatal(err)
	}

	data.Scan()
	if record, err = strconv.Atoi(parseValues(data.Text())); err != nil {
		log.Fatal(err)
	}

	totalMargin := 0
	for holdTime := 0; holdTime < time; holdTime++ {
		speed := holdTime

		if (time-holdTime)*speed > record {
			totalMargin++
		}
	}

	total *= totalMargin

	fmt.Printf("Part1 total: %d\n", total)
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

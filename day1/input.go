package main

import (
	"bufio"
	"os"
)

type InputFile string

const (
	InputFileExample1 InputFile = "example1.txt"
	InputFileExample2 InputFile = "example2.txt"
	InputFileReal1    InputFile = "input1.txt"
	InputFileReal2    InputFile = "input2.txt"
)

func readInput(filename InputFile) (*bufio.Scanner, error) {
	file, err := os.Open(string(filename))
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

package iolib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

// ReadInt parses next line of stdin to int.
func ReadInt() int {
	if ok := scanner.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading from stdin: %v\n", err)
	}

	input := scanner.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("unable to parse int from input: %v\n", err)
	}

	return n
}

// ReadInts parses next line of stdin to int slice.
func ReadInts() []int {
	if ok := scanner.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading from stdin: %v\n", err)
	}

	input := strings.Split(scanner.Text(), " ")
	intSlice := make([]int, len(input))
	for i := 0; i < len(intSlice); i++ {
		n, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatalf("unable to parse int from input: %v\n", err)
		}
		intSlice[i] = n
	}

	return intSlice
}

// ReadString parses next line of stdin to string.
func ReadString() string {
	if ok := scanner.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading from stdin: %v\n", err)
	}

	return scanner.Text()
}

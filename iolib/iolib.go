package iolib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInt() int {
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	input := s.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("unable to parse int from input: %v\n", err)
	}

	return n
}

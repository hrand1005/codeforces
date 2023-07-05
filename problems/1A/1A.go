package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(n, m, a int) {
	perRow := n / a
	if n%a != 0 {
		perRow++
	}
	perCol := m / a
	if m%a != 0 {
		perCol++
	}
	fmt.Println(perRow * perCol)
}

// ReadInts parses next line of stdin to int slice.
func ReadInts() []int {
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	if err := s.Err(); err != nil {
		log.Fatalf("error reading from stdin: %v\n", err)
	}

	input := strings.Split(s.Text(), " ")
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

func main() {
	ints := ReadInts()
	solve(ints[0], ints[1], ints[2])
}

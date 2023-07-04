package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	input := s.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("unable to parse int from input: %v\n", err)
	}

	if n > 2 && n%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

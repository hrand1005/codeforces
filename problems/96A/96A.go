package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solve(s string) {
	consecutive := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			consecutive++
		} else {
			consecutive = 1
		}
		if consecutive == 7 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

// ReadString parses next line of stdin to string.
func ReadString() string {
	s := bufio.NewScanner(os.Stdin)
	if ok := s.Scan(); !ok {
		log.Fatal("failed to read from stdin")
	}

	if err := s.Err(); err != nil {
		log.Fatalf("error reading from stdin: %v\n", err)
	}

	return s.Text()
}

func main() {
	s := ReadString()
	solve(s)
}

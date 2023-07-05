package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func solve(s string) string {
	res := ""
	vowels := "AEIOUYaeiouy"
	for _, v := range s {
		if unicode.IsLetter(v) && !strings.Contains(vowels, string(v)) {
			res += fmt.Sprintf(".%s", string(unicode.ToLower(v)))
		}
	}
	return res
}

func readString(r *bufio.Reader) string {
	s, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read string: %v\n", err)
	}
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s := readString(r)
	result := solve(s)
	w.WriteString(fmt.Sprintf("%v\n", result))
}

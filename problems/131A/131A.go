package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isCapsLock(s string) bool {
	for _, l := range s[1:] {
		if !unicode.IsUpper(l) {
			return false
		}
	}
	return true
}

func solve(s string) string {
	if !isCapsLock(s) {
		return s
	}
	sNew := strings.Builder{}
	for _, v := range s {
		if unicode.IsUpper(v) {
			sNew.WriteRune(unicode.ToLower(v))
		} else {
			sNew.WriteRune(unicode.ToUpper(v))
		}
	}
	return sNew.String()
}

func stringToIntSlice(s string) []int {
	nums := strings.Split(s, " ")
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to parse int: %v\n", err)
		}
		ints[i] = n
	}
	return ints
}

func stringSliceToIntSlice(nums []string) []int {
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to parse int: %v\n", err)
		}
		ints[i] = n
	}
	return ints
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

func readInt(r *bufio.Reader) int {
	s := readString(r)
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v\n", err)
	}
	return n
}

func readStringSlice(r *bufio.Reader) []string {
	s := readString(r)
	return strings.Split(s, " ")
}

func readIntSlice(r *bufio.Reader) []int {
	nums := readStringSlice(r)
	return stringSliceToIntSlice(nums)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	result := solve(readString(r))
	w.WriteString(fmt.Sprintf("%v\n", result))
}

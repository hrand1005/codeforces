package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(k int, plankHeights []int) int {
	minHeight := 0
	for i := 0; i < k; i++ {
		minHeight += plankHeights[i]
	}

	height := minHeight
	minIdx := 0
	for i := k; i < len(plankHeights); i++ {
		height = height - plankHeights[i-k] + plankHeights[i]
		if height < minHeight {
			minHeight = height
			minIdx = i - k + 1
		}
	}
	// NOTE: one indexed
	return minIdx + 1
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

	nk := readIntSlice(r)
	k := nk[1]
	plankHeights := readIntSlice(r)

	result := solve(k, plankHeights)
	w.WriteString(fmt.Sprintf("%v\n", result))
}

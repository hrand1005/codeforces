package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(prices, money []int) []int {
	const shopMax = 100000
	available := make([]int, shopMax+1)
	for _, p := range prices {
		available[p]++
	}

	for i := 1; i < len(available); i++ {
		available[i] += available[i-1]
	}

	buy := make([]int, len(money))
	for i, m := range money {
		if m > shopMax {
			buy[i] = len(prices)
		} else {
			buy[i] = available[m]
		}
	}
	return buy
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

	_ = readInt(r)
	prices := readIntSlice(r)
	days := readInt(r)
	money := make([]int, days)
	for i := 0; i < days; i++ {
		money[i] = readInt(r)
	}

	result := solve(prices, money)
	for _, r := range result {
		w.WriteString(fmt.Sprintf("%v\n", r))
	}
}

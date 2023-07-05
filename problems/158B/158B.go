package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getCarGroup(groups []int) []int {
	count := 0
	group := make([]int, 0)
	for i := len(groups) - 1; i >= 0; i-- {
		if count == 4 || count+groups[0] > 4 {
			break
		}
		if count+groups[i] <= 4 {
			group = append(group, i)
			count += groups[i]
		}
	}
	return group
}

func popCarGroup(cg, groups []int) []int {
	for _, v := range cg {
		groups = append(groups[:v], groups[v+1:]...)
	}
	return groups
}

func solve(n int, groups []int) int {
	sort.Ints(groups)

	cars := 0
	for len(groups) != 0 {
		carGroup := getCarGroup(groups)
		cars++
		groups = popCarGroup(carGroup, groups)
	}
	return cars
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

	n := readInt(r)
	groups := readIntSlice(r)
	result := solve(n, groups)
	w.WriteString(fmt.Sprintf("%v\n", result))
}

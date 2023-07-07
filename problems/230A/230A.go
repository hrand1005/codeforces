package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dragon struct {
	strength int
	bonus    int
}

func solve(s int, dragons []dragon) string {
	for len(dragons) != 0 {
		dIdx := -1
		for i, d := range dragons {
			if d.strength < s {
				dIdx = i
				break
			}
		}
		if dIdx == -1 {
			return "NO"
		}
		this := dragons[dIdx]
		dragons = append(dragons[:dIdx], dragons[dIdx+1:]...)
		s += this.bonus
	}
	return "YES"
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

	strengthAndCount := readIntSlice(r)
	strength := strengthAndCount[0]
	count := strengthAndCount[1]
	dragons := make([]dragon, count)
	for i := 0; i < count; i++ {
		strengthAndBonus := readIntSlice(r)
		dragons[i] = dragon{
			strength: strengthAndBonus[0],
			bonus:    strengthAndBonus[1],
		}
	}

	result := solve(strength, dragons)
	w.WriteString(fmt.Sprintf("%v\n", result))
}

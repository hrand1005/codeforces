package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Manually typed out version:
//
// var luckyNumbers = []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 477, 744, 747, 774, 777}
// func slove(n int) string {
// 	for _, l := range stringSliceToIntSlice(luckyNumbers) {
// 		if n%l == 0 {
// 			return "YES"
// 		}
// 	}
// 	return "NO"
// }

func generateLuckyNumbers(luckyNumbers []string, n int) []string {
	if n == 0 {
		return nil
	}
	newNumbers := []string{}
	for _, v := range luckyNumbers {
		newNumbers = append(newNumbers, v+"4")
		newNumbers = append(newNumbers, v+"7")
	}
	newNumbers = append(luckyNumbers, newNumbers...)
	return append(newNumbers, generateLuckyNumbers(newNumbers, n-1)...)
}

func solve(n int) string {
	luckyNumbers := generateLuckyNumbers([]string{"4", "7"}, 2)
	fmt.Printf("%#v\n", luckyNumbers)
	for _, l := range stringSliceToIntSlice(luckyNumbers) {
		if n%l == 0 {
			return "YES"
		}
	}
	return "NO"
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

	result := solve(readInt(r))
	w.WriteString(fmt.Sprintf("%v\n", result))
}

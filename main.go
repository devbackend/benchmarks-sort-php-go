package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
}

func SortFile(in string) {
	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func QuickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	c := input[0]

	var left, right []int

	for _, v := range input[1:] {
		if v < c {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	var result []int = []int{}

	result = append(result, QuickSort(left)...)
	result = append(result, c)
	result = append(result, QuickSort(right)...)

	return result
}

func Compare(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}

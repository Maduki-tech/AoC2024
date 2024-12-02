package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/inputDay1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// get all chars thill space
		lines = append(lines, scanner.Text())
	}

	var first, last []int

	// example string: "3   4"
	for i := 0; i < len(lines); i++ {
		// get first and last number
		left, err := strconv.Atoi(strings.Split(lines[i], " ")[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(strings.Split(lines[i], " ")[3])
		if err != nil {
			panic(err)
		}

		// add to array
		first = append(first, left)
		last = append(last, right)
	}

	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})

	sort.Slice(last, func(i, j int) bool {
		return last[i] < last[j]
	})

	part1(first, last)

	part2(first, last)
}

func part1(first []int, last []int) {
	var result int
	for i := 0; i < len(first); i++ {
		if last[i] < first[i] {
			result += first[i] - last[i]
		} else {
			result += last[i] - first[i]
		}
	}

	fmt.Println(result)

}

func part2(first []int, last []int) {
	var result int
	var counter int
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(last); j++ {
			if first[i] == last[j] {
				fmt.Println(first[i], last[j])
				counter++
			}
		}

		result += first[i] * counter
		counter = 0
	}

	fmt.Println(result)

}

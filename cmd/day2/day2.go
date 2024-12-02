package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("cmd/day2/inputDay2.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var numbers []int

	var levels int
	for i := 0; i < len(lines); i++ {

		line := strings.Split(lines[i], " ")
		levels = len(line)

		for j := 0; j < len(line); j++ {
			res, err := strconv.Atoi(line[j])
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, res)
		}
	}

	res := part1(numbers, levels)
	fmt.Println(res)
	// part2(numbers)
}

func part1(numbers []int, levels int) int {
	var result int
	for i := 0; i < len(numbers); i++ {
		if i%levels == 0 {
			fmt.Println()
		}
	}

	return result
}

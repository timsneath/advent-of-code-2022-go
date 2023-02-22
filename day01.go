package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countElfCalories(lines []string) []int {
	var elfCalories []int
	var currentCalories int

	for _, row := range lines {

		if row != "" {
			c, _ := strconv.Atoi(row)
			currentCalories += c
		} else {
			elfCalories = append(elfCalories, currentCalories)
			currentCalories = 0
		}
	}
	return elfCalories
}

func sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func part1(lines []string) {
	calories := countElfCalories(lines)
	sort.Ints(calories)
	max := calories[len(calories)-1]

	fmt.Printf("Part 1: %d\n", max)
}

func part2(lines []string) {
	calories := countElfCalories(lines)
	sort.Ints(calories)
	top3 := calories[len(calories)-3:]
	sumTop3 := sum(top3)

	fmt.Printf("Part 2: %d\n", sumTop3)
}

func main() {
	data, err := os.ReadFile("data/day01.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "day01: %v\n", err)
		os.Exit(-1)
	}
	lines := strings.Split(string(data), "\n")

	part1(lines)
	part2(lines)
}

package main

import (
	//"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// Process the input into two sorted arrays
	lines := strings.Split(input, "\n")

	var list1 []int
	var list2 []int

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		if value1, err := strconv.Atoi(parts[0]); err == nil {
			list1 = append(list1, value1)
		}
		if value2, err := strconv.Atoi(parts[1]); err == nil {
			list2 = append(list2, value2)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	// calculate the similarity score and return the result for part 2
	if part2 {
		dictList2 := make(map[int]int)
		for _, num := range list2 {
			dictList2[num] = dictList2[num] + 1
		}

		var similarity int
		similarity = 0
		for lcv := 0; lcv < len(list1); lcv++ {
			similarity = similarity + (list1[lcv] * dictList2[list1[lcv]])
		}
		return similarity
	}

	// calculate the distance between the two and return the result for part 1
	var distance int
	distance = 0
	for lcv := 0; lcv < len(list1); lcv++ {
		distance = distance + int(math.Abs(float64(list1[lcv]-list2[lcv])))
	}

	return distance
}

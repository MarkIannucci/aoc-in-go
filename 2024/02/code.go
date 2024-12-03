package main

import (
	"strings"
	"slices"
	"fmt"
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

	var safeLines int
	safeLines = 0
	for _, line := range lines {
		parts := strings.Split(line, " ")

		// convert the parts to integers
		partsInt := make([]int, len(parts))
		for i, part := range parts {
			if value, err := strconv.Atoi(part); err == nil {
				partsInt[i] = value
			}
		
		sorted := make([]int, len(partsInt))
		copy(sorted, partsInt)
		slices.Sort(sorted)

		if slices.Equal(sorted, parts) {
			// do nothing because we need to continue evaluating if it is potentially safe
		} else {
			slices.Reverse(sorted)
			if slices.Equal(sorted, parts) {
				slices.Sort(sorted)
			} else {
				break
			}
		}

		// check if each number in the sorted string is increasing by at least 1 and not more than 3	
		for i := 0; i < len(sorted) - 1; i++ {
			if sorted[i + 1] - sorted[i] > 3 {
				isSafe = false
				break
			}
		}

	}

	
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return 43
}
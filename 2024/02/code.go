package main

import (
	"fmt"
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
		}

		sorted := make([]int, len(partsInt))
		copy(sorted, partsInt)
		slices.Sort(sorted)

		if slices.Equal(sorted, partsInt) {
			// do nothing because we need to continue evaluating if it is potentially safe
		} else {
			slices.Reverse(sorted)
			if slices.Equal(sorted, partsInt) {
				slices.Sort(sorted)
			} else {
				continue
			}
		}

		// check if each number in the sorted string is increasing by at least 1 and not more than 3
		var safe bool
		safe = true
		for i := 0; i < len(sorted)-1; i++ {
			if sorted[i+1]-sorted[i] == 0 {
				safe = false
			}
			if sorted[i+1]-sorted[i] >= 4 {
				safe = false
			}
		}
		if safe {
			safeLines = safeLines + 1
		}

	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
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
			}

			var ruleViolators int
			ruleViolators = 0
			// determine if the parts are increasing or decreasion
			fmt.Println(partsInt)
			if partsInt[0] < partsInt[len(partsInt)-1] {
				
				// increasing
				// itreate through the parts and check if the difference between each part is at least 1 and not more than 3
				for i := 0; i < len(partsInt)-1; i++ {
					if partsInt[i+1]-partsInt[i] <= 0 {
						ruleViolators = ruleViolators + 1
					}
					if partsInt[i+1]-partsInt[i] >= 4 {
						ruleViolators = ruleViolators + 1
					}
				}
			} else {
				for i := 0; i < len(partsInt)-1; i++ {
					if partsInt[i+1]-partsInt[i] >= 0 {
						ruleViolators = ruleViolators + 1
					}
					if partsInt[i+1]-partsInt[i] <= -4 {
						ruleViolators = ruleViolators + 1
					}
				}
			}
			fmt.Println(ruleViolators)
			
			if ruleViolators < 2 {
				safeLines = safeLines + 1
			}
		}
		return safeLines
	}
	// solve part 1 here
	return safeLines
}

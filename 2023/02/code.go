package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile("(\\d+) (red|green|blue)")
	ans := 0

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, line := range strings.Split(input, "\n") {
			maxRed := 0
			maxGreen := 0
			maxBlue := 0
			for _, matches := range re.FindAllStringSubmatch(line, -1) {
				number, _ := strconv.Atoi(matches[1])
				color := matches[2]

				switch color {
				case "red":
					maxRed = max(maxRed, number)
				case "green":
					maxGreen = max(maxGreen, number)
				case "blue":
					maxBlue = max(maxBlue, number)
				}
			}
			ans += maxRed * maxGreen * maxBlue
		}
		return ans
	}
	// solve part 1 here
	for gameId, line := range strings.Split(input, "\n") {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, matches := range re.FindAllStringSubmatch(line, -1) {
			number, _ := strconv.Atoi(matches[1])
			color := matches[2]

			switch color {
			case "red":
				maxRed = max(maxRed, number)
			case "green":
				maxGreen = max(maxGreen, number)
			case "blue":
				maxBlue = max(maxBlue, number)
			}
		}
		if maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14 {
			ans += gameId + 1
		}
	}
	return ans
}

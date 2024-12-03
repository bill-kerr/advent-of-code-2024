package day02

import (
	"fmt"
	"strings"

	"github.com/bill-kerr/advent-of-code-2024/util"
)

func part1(lines []string) {
	safeReports := 0

	for _, line := range lines {
		values := util.ParseInts(strings.Split(line, " "))
		if checkSafety(values, -1) {
			safeReports++
		}
	}

	fmt.Printf("part1: %d\n", safeReports)
}

func part2(lines []string) {
	safeReports := 0

	for _, line := range lines {
		values := util.ParseInts(strings.Split(line, " "))

		for i := -1; i < len(values); i++ {
			if isSafe := checkSafety(values, i); isSafe {
				safeReports++
				break
			}
		}
	}

	fmt.Printf("part2: %d\n", safeReports)
}

// a skipIndex of < 0 indicates that no values should be skipped
func checkSafety(values []int, skipIndex int) bool {
	isIncreasing := getIsIncreasing(values, skipIndex)
	lastValue := values[0]
	startIndex := 0

	if skipIndex == 0 {
		lastValue = values[1]
		startIndex = 1
	}

	for i, value := range values {
		if i == skipIndex || i == startIndex {
			continue
		}

		if value == lastValue {
			return false
		}

		if isIncreasing && value < lastValue {
			return false
		}

		if !isIncreasing && value > lastValue {
			return false
		}

		if util.AbsInt(value-lastValue) > 3 {
			return false
		}

		lastValue = value
	}

	return true
}

func getIsIncreasing(values []int, skipIndex int) bool {
	firstVal := values[0]
	secondVal := values[1]

	if skipIndex == 0 {
		firstVal = values[1]
		secondVal = values[2]
	} else if skipIndex == 1 {
		secondVal = values[2]
	}

	return secondVal > firstVal
}

func Run() {
	lines := util.OpenAndRead("./day02/input.txt")

	part1(lines)
	part2(lines)
}

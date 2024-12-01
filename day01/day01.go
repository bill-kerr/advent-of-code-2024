package day01

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bill-kerr/advent-of-code-2024/util"
)

func part1(lines []string) {
	leftNums := make([]int, len(lines))
	rightNums := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		leftNums[i] = util.Atoi(parts[0])
		rightNums[i] = util.Atoi(parts[1])
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	sumOfDistances := 0
	for i, leftNum := range leftNums {
		rightNum := rightNums[i]
		sumOfDistances += util.AbsInt(leftNum - rightNum)
	}

	fmt.Printf("part 1: %d\n", sumOfDistances)
}

func part2(lines []string) {
	leftNums := make([]int, len(lines))
	freqMap := make(map[int]int)

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		leftNums[i] = util.Atoi(parts[0])
		rightNum := util.Atoi(parts[1])
		freqMap[rightNum]++
	}

	sumOfScores := 0
	for _, num := range leftNums {
		sumOfScores += num * freqMap[num]
	}

	fmt.Printf("part 2: %d\n", sumOfScores)
}

func Run() {
	lines := util.OpenAndRead("./day01/input.txt")

	part1(lines)
	part2(lines)
}

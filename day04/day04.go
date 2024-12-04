package day04

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2024/util"
)

func Run() {
	lines := util.OpenAndRead("./day04/input.txt")

	part1(lines)
	part2(lines)
}

const xmas = "XMAS"

type direction string

const (
	left      direction = "left"
	right     direction = "right"
	up        direction = "up"
	down      direction = "down"
	upLeft    direction = "upLeft"
	upRight   direction = "upRight"
	downLeft  direction = "downLeft"
	downRight direction = "downRight"
)

func part1(lines []string) {
	count := 0
	slicer := &slicer{data: lines, wordLength: len(xmas)}

	for row, line := range lines {
		for col, char := range line {
			if char != 'X' {
				continue
			}

			canLeft := col > 2
			canRight := col < len(line)-3
			canUp := row >= 3
			canDown := row < len(lines)-3

			// left
			if canLeft && slicer.slice(left, row, col) == xmas {
				count++
			}

			// right
			if canRight && slicer.slice(right, row, col) == xmas {
				count++
			}

			// up
			if canUp && slicer.slice(up, row, col) == xmas {
				count++
			}

			// down
			if canDown && slicer.slice(down, row, col) == xmas {
				count++
			}

			// up-left
			if canLeft && canUp && slicer.slice(upLeft, row, col) == xmas {
				count++
			}

			// up-right
			if canRight && canUp && slicer.slice(upRight, row, col) == xmas {
				count++
			}

			// down-left
			if canLeft && canDown && slicer.slice(downLeft, row, col) == xmas {
				count++
			}

			// down-right
			if canRight && canDown && slicer.slice(downRight, row, col) == xmas {
				count++
			}
		}
	}

	fmt.Printf("part1: %d\n", count)
}

type slicer struct {
	data       []string
	wordLength int
}

func (s *slicer) slice(direction direction, row, col int) string {
	var word string

	switch direction {
	case left:
		word = s.getWord(stay(row), decrease(col))
	case right:
		word = s.getWord(stay(row), increase(col))
	case up:
		word = s.getWord(decrease(row), stay(col))
	case down:
		word = s.getWord(increase(row), stay(col))
	case upLeft:
		word = s.getWord(decrease(row), decrease(col))
	case upRight:
		word = s.getWord(decrease(row), increase(col))
	case downLeft:
		word = s.getWord(increase(row), decrease(col))
	case downRight:
		word = s.getWord(increase(row), increase(col))
	default:
	}

	return word
}

func (s *slicer) getWord(incrementRow, incrementCol func(int) int) string {
	var word string

	for i := range s.wordLength {
		word += string(s.data[incrementRow(i)][incrementCol(i)])
	}

	return word
}

func stay(rowOrCol int) func(int) int {
	return func(iter int) int {
		return rowOrCol
	}
}

func increase(rowOrCol int) func(int) int {
	return func(iter int) int {
		return rowOrCol + iter
	}
}

func decrease(rowOrCol int) func(int) int {
	return func(iter int) int {
		return rowOrCol - iter
	}
}

func part2(lines []string) {
	count := 0
	slicer := &slicer{data: lines, wordLength: 3}

	for row, line := range lines {
		for col, char := range line {
			if char != 'A' || row == 0 || col == len(lines[0])-1 || row == len(lines)-1 || col == 0 {
				continue
			}

			topLeft := slicer.slice(downRight, row-1, col-1)
			topRight := slicer.slice(downLeft, row-1, col+1)

			if (topLeft == "MAS" || topLeft == "SAM") && (topRight == "MAS" || topRight == "SAM") {
				count++
			}
		}
	}

	fmt.Printf("part2: %d\n", count)
}

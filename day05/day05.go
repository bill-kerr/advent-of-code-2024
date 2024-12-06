package day05

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/bill-kerr/advent-of-code-2024/util"
)

func part1(lines []string) {
	rules, printings := parseInput(lines)
	sum := 0

	for _, printing := range printings {
		if printingIsValid(printing, rules) {
			// make sure the length of the printing is an odd number
			if len(printing)%2 == 0 {
				panic("printing not an odd length")
			}

			sum += printing[len(printing)/2]
		}
	}

	fmt.Printf("part1: %d\n", sum)
}

func part2(lines []string) {
	rules, printings := parseInput(lines)
	sum := 0

	for _, printing := range printings {
		isValid, fixedPrinting := maybeFixPrinting(printing, rules)
		if !isValid {
			// make sure the length of the printing is an odd number
			if len(fixedPrinting)%2 == 0 {
				panic("printing not an odd length")
			}

			if !printingIsValid(fixedPrinting, rules) {
				log.Fatalf("fixed printing not valid: %v\n", fixedPrinting)
			}

			sum += fixedPrinting[len(printing)/2]
		}
	}

	fmt.Printf("part1: %d\n", sum)
}

func printingIsValid(printing []int, rules map[int]map[int]struct{}) bool {
	seenPages := make(map[int]struct{})

	for _, pageNumber := range printing {
		seenPages[pageNumber] = struct{}{}

		if rule, ok := rules[pageNumber]; ok {
			for priorPageNumber := range rule {
				printingHasPage := printingIncludes(printing, priorPageNumber)
				_, hasSeenPage := seenPages[priorPageNumber]

				if printingHasPage && !hasSeenPage {
					return false
				}
			}
		}
	}

	return true
}

func maybeFixPrinting(printing []int, rules map[int]map[int]struct{}) (bool, []int) {
	seenPages := make(map[int]struct{})
	priors := make(map[int]int)
	valid := true

	for _, pageNumber := range printing {
		priors[pageNumber] = 0
		seenPages[pageNumber] = struct{}{}

		if rule, ok := rules[pageNumber]; ok {
			for priorPageNumber := range rule {
				priorPageIndex := slices.Index(printing, priorPageNumber)
				_, hasSeenPage := seenPages[priorPageNumber]

				if priorPageIndex > -1 && !hasSeenPage {
					valid = false
				}

				if priorPageIndex > -1 {
					priors[pageNumber]++
				}
			}
		}
	}

	slices.SortFunc(printing, func(a, b int) int {
		return priors[a] - priors[b]
	})

	return valid, printing
}

func printingIncludes(printing []int, pageNumber int) bool {
	for _, pageNum := range printing {
		if pageNumber == pageNum {
			return true
		}
	}
	return false
}

func parseInput(lines []string) (map[int]map[int]struct{}, [][]int) {
	parsingRules := true
	rules := map[int]map[int]struct{}{}
	printings := [][]int{}

	for _, line := range lines {
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			parts := util.ParseInts(strings.Split(line, "|"))

			if _, ok := rules[parts[1]]; !ok {
				rules[parts[1]] = make(map[int]struct{})
			}
			rules[parts[1]][parts[0]] = struct{}{}
		} else {
			pageNumbers := util.ParseInts(strings.Split(line, ","))
			printings = append(printings, pageNumbers)
		}
	}

	return rules, printings
}

func Run() {
	lines := util.OpenAndRead("./day05/input.txt")

	part1(lines)
	part2(lines)
}

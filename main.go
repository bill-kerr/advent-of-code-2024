package main

import (
	"os"

	"github.com/bill-kerr/advent-of-code-2024/day04"
	"github.com/bill-kerr/advent-of-code-2024/util"
)

func main() {
	if len(os.Args) >= 3 {
		t := os.Args[1]
		name := os.Args[2]

		if t == "-t" && name != "" {
			util.CreateTemplate(name)
		}
	} else {
		day04.Run()
	}
}

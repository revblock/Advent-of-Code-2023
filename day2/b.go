package day2

import (
	"fmt"

	"github.com/m/v2/revblock/Advent-of-Code-2023/input"
)

func B() {

	scanner, file := input.NewFileScanner("./data/2.txt")
	defer file.Close()

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		_, maxCubes := parseLine(line)

		power := maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
		result = result + power
	}

	fmt.Println(result)
}

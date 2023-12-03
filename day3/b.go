package day3

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func B() {
	data, _ := os.ReadFile("./data/3.txt")
	input := string(data)

	lineWidth := getLineWidth(input)
	runes := toRunes(input)

	var gears []int
	numbers := make(map[int]*[]rune)
	var currentNumber *[]rune
	newNumber := make([]rune, 0)
	currentNumber = &newNumber

	for i, r := range runes {
		// If value is a number, save value and map pointer
		if slices.Contains(numberRunes, r) {
			*currentNumber = append(*currentNumber, r)
			numbers[i] = currentNumber
		}

		if !slices.Contains(numberRunes, r) {
			newNumber := make([]rune, 0)
			currentNumber = &newNumber

			if r == '*' {
				gears = append(gears, i)
			}
		}
	}
	result := 0

	for _, g := range gears {
		adjacentIndexes := getAdjacentIndexes(g, lineWidth)

		var adjacentNumbers []*[]rune

		for _, ai := range adjacentIndexes {
			if number, ok := numbers[ai]; ok && !slices.Contains(adjacentNumbers, number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}

		if len(adjacentNumbers) == 2 {
			num1, _ := strconv.Atoi(string(*adjacentNumbers[0]))
			num2, _ := strconv.Atoi(string(*adjacentNumbers[1]))
			result = result + (num1 * num2)
		}
	}

	fmt.Println(result)

}

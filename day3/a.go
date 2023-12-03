package day3

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

var numberRunes = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func A() {
	data, _ := os.ReadFile("./data/3.txt")
	input := string(data)

	lineWidth := getLineWidth(input)
	runes := toRunes(input)

	var adjacentIndexes []int
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

			if r != '.' {
				adjacentIndexes = append(adjacentIndexes, getAdjacentIndexes(i, lineWidth)...)
			}
		}
	}

	var adjacentNumbers []*[]rune
	results := 0

	for _, ai := range adjacentIndexes {
		if number, ok := numbers[ai]; ok && !slices.Contains(adjacentNumbers, number) {
			adjacentNumbers = append(adjacentNumbers, number)
			intValue, _ := strconv.Atoi(string(*number))
			results = results + intValue
		}
	}

	fmt.Println(results)

}

func getAdjacentIndexes(n int, lineWidth int) []int {
	left := n - 1
	right := n + 1
	topLeft := left - lineWidth
	top := n - lineWidth
	topRight := right - lineWidth
	bottomLeft := left + lineWidth
	bottom := n + lineWidth
	bottomRight := right + lineWidth

	return []int{topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight}
}

func getLineWidth(s string) int {
	for i, c := range s {
		if c == '\n' {
			return i
		}
	}
	return len(s)
}

func toRunes(s string) (runes []rune) {
	for _, r := range []rune(s) {
		if r != '\n' {
			runes = append(runes, r)
		}
	}

	return
}

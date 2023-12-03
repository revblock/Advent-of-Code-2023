package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/m/v2/revblock/Advent-of-Code-2023/input"
)

var numberValues = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

var wordValues = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func B() {

	scanner, file := input.NewFileReader("./data/1.txt")
	defer file.Close()

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		tokens := tokenize(line, 5)

		var firstDigit string
		var lastDigit string
		for i := 0; i < len(tokens); i++ {

			var digit string

			if numberIndex := slices.Index(numberValues, tokens[i]); numberIndex != -1 {
				digit = numberValues[numberIndex]
			}

			if digit == "" {
				if wordIndex := slices.Index(wordValues, tokens[i]); wordIndex != -1 {
					digit = numberValues[wordIndex+1]
				}
			}

			if digit != "" {
				if firstDigit == "" {
					firstDigit = digit
				}

				lastDigit = digit
			}
		}

		value, _ := strconv.Atoi(firstDigit + lastDigit)

		result = result + value
	}

	fmt.Println(result)
}

func tokenize(s string, maxTokenLength int) []string {

	chars := strings.Split(s, "")

	var tokens []string

	for startIndex := 0; startIndex < len(chars); startIndex++ {

		for endIndex := 1; startIndex+endIndex <= len(chars) && endIndex <= maxTokenLength; endIndex++ {

			var comparison string

			if endIndex == 1 {
				comparison = chars[startIndex]
			} else {
				comparison = strings.Join(chars[startIndex:startIndex+endIndex], "")
			}

			tokens = append(tokens, comparison)
		}
	}

	return tokens
}

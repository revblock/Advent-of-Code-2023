package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/m/v2/revblock/Advent-of-Code-2023/input"
)

func A() {

	scanner, file := input.NewFileScanner("./data/1.txt")
	defer file.Close()

	// Any digit character
	r, _ := regexp.Compile("\\d")

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Match all digits
		matches := r.FindAllString(line, -1)

		// Get first and last
		first := matches[0]
		last := matches[len(matches)-1]

		// Concat and convert to int
		value, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

		result = result + value
	}

	fmt.Println(result)
}

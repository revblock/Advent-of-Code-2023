package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/m/v2/revblock/Advent-of-Code-2023/input"
)

func A() {

	scanner, file := input.NewFileReader("./data/1.txt")
	defer file.Close()

	r, _ := regexp.Compile("\\d")

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllString(line, -1)
		first := matches[0]
		last := matches[len(matches)-1]

		value, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

		result = result + value
	}

	fmt.Println(result)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m/v2/revblock/Advent-of-Code-2023/day1"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Day and Part:")
	dayPart, _ := reader.ReadString('\n')

	if strings.TrimSuffix(dayPart, "\n") == "1A" {
		day1.A()
	}

	if strings.TrimSuffix(dayPart, "\n") == "1B" {
		day1.B()
	}
}

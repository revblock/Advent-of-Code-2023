package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/m/v2/revblock/Advent-of-Code-2023/input"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func A() {

	scanner, file := input.NewFileScanner("./data/2.txt")
	defer file.Close()

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameID, maxCubes := parseLine(line)

		if maxCubes["red"] <= maxRed && maxCubes["green"] <= maxGreen && maxCubes["blue"] <= maxBlue {
			result = result + gameID
		}
	}

	fmt.Println(result)
}

func parseLine(s string) (gameID int, maxCubes map[string]int) {

	gameParts := trimArray(strings.Split(s, ":"))
	gameIDParts := trimArray(strings.Split(gameParts[0], " "))
	gameID, _ = strconv.Atoi(gameIDParts[1])
	roundParts := trimArray(strings.Split(gameParts[1], ";"))
	maxCubes = make(map[string]int)
	for _, rp := range roundParts {
		for _, cp := range trimArray(strings.Split(rp, ",")) {
			cube := trimArray(strings.Split(cp, " "))
			amount, _ := strconv.Atoi(cube[0])
			colour := cube[1]
			cubeValue, exists := maxCubes[colour]

			if (exists && amount > cubeValue) || !exists {
				maxCubes[colour] = amount
			}
		}
	}

	return
}

func trimArray(a []string) (trimmed []string) {
	for _, s := range a {
		trimmed = append(trimmed, strings.TrimSpace(s))
	}
	return
}

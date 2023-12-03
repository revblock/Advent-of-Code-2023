package input

import (
	"bufio"
	"os"
)

func NewFileScanner(filename string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	return scanner, file
}

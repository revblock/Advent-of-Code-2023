package input

import (
	"bufio"
	"os"
)

func NewFileReader(filename string) (*bufio.Scanner, *os.File) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	return scanner, file
}

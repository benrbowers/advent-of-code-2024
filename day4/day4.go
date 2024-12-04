package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Looking for "XMAS" in word search.

It can be backwards/forwards, vertical/horizontal,
diagonal, or overlapping other occurrances
*/

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input file:")
		panic(err)
	}
	defer inputFile.Close()

	var horizontal []string = []string{}
	var vertical []string = []string{}
	var upDiagonal []string = []string{}
	var downDiagonal []string = []string{}

	scanner := bufio.NewScanner(inputFile)

	// Collect horizontal lines
	for scanner.Scan() {
		line := scanner.Text()
		horizontal = append(horizontal, line)
	}

	var lineCount int = len(horizontal)
	var lineLength int = len(horizontal[0])

	// Collect vertical lines
	for i := range lineLength {
		var column []byte = []byte{}

		for _, row := range horizontal {
			column = append(column, row[i])
		}
		vertical = append(vertical, string(column))
	}

	// Collect up diagonal lines
	for row := range lineCount {
		var diag []byte = []byte{}

		for col := range min(row+1, lineLength) {
			diag = append(diag, horizontal[row-col][col])
		}
		upDiagonal = append(upDiagonal, string(diag))
	}
	for col := 1; col < lineLength; col++ {
		var diag []byte = []byte{}

		for delta := range min(lineLength-col, lineCount) {
			row := lineCount - 1 - delta
			diag = append(diag, horizontal[row][col+delta])
		}
		upDiagonal = append(upDiagonal, string(diag))
	}

	// Collect down diagonal lines
	for row := lineCount - 1; row >= 0; row-- {
		var diag []byte = []byte{}

		for col := range min(lineCount-row, lineLength) {
			diag = append(diag, horizontal[row+col][col])
		}
		downDiagonal = append(downDiagonal, string(diag))
	}
	for col := 1; col < lineLength; col++ {
		var diag []byte = []byte{}

		for row := range min(lineLength-col, lineCount) {
			diag = append(diag, horizontal[row][col+row])
		}
		downDiagonal = append(downDiagonal, string(diag))
	}

	var answer int = 0
	const forwards string = "XMAS"
	const backwards string = "SAMX"

	for _, line := range horizontal {
		answer += strings.Count(line, forwards)
		answer += strings.Count(line, backwards)
	}
	for _, line := range vertical {
		answer += strings.Count(line, forwards)
		answer += strings.Count(line, backwards)
	}
	for _, line := range upDiagonal {
		answer += strings.Count(line, forwards)
		answer += strings.Count(line, backwards)
	}
	for _, line := range downDiagonal {
		answer += strings.Count(line, forwards)
		answer += strings.Count(line, backwards)
	}

	fmt.Printf("Occurences of \"XMAS\": %d\n", answer)
}

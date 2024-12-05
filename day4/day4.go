package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	for col := lineLength - 1; col >= 0; col-- {
		var diag []byte = []byte{}

		for row := range min(lineLength-col, lineCount) {
			diag = append(diag, horizontal[row][col+row])
		}
		downDiagonal = append(downDiagonal, string(diag))
	}
	for row := 1; row < lineLength; row++ {
		var diag []byte = []byte{}

		for col := range min(lineCount-row, lineLength) {
			diag = append(diag, horizontal[row+col][col])
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

	masPattern := regexp.MustCompile("MAS")
	samPattern := regexp.MustCompile("SAM")

	// Coords for a single 'A' in an x-MAS
	var upMatches [][]int = [][]int{}
	var downMatches [][]int = [][]int{}

	for i, line := range upDiagonal {
		for _, indexes := range masPattern.FindAllStringIndex(line, -1) {
			position := indexes[0] + 1

			var row int
			var col int

			if i < lineCount {
				row = i - position
				col = position
			} else {
				row = lineCount - 1 - position
				col = i - lineCount + 1
			}

			upMatches = append(upMatches, []int{row, col})
		}
		for _, indexes := range samPattern.FindAllStringIndex(line, -1) {
			position := indexes[0] + 1

			var row int
			var col int

			if i < lineCount {
				row = i - position
				col = position
			} else {
				row = lineCount - 1 - position
				col = i - lineCount + 1
			}

			upMatches = append(upMatches, []int{row, col})
		}
	}

	for i, line := range downDiagonal {
		for _, indexes := range masPattern.FindAllStringIndex(line, -1) {
			position := indexes[0] + 1

			var row int
			var col int

			if i < lineLength {
				row = position
				col = lineLength - 1 - i + position
			} else {
				row = (i - lineCount) + 1 + position
				col = position
			}

			downMatches = append(downMatches, []int{row, col})
		}
		for _, indexes := range samPattern.FindAllStringIndex(line, -1) {
			position := indexes[0] + 1

			var row int
			var col int

			if i < lineLength {
				row = position
				col = lineLength - 1 - i + position
			} else {
				row = (i - lineCount) + 1 + position
				col = position
			}

			downMatches = append(downMatches, []int{row, col})
		}
	}

	var xMasCount int = 0

	for _, upMatch := range upMatches {
		for _, downMatch := range downMatches {
			if upMatch[0] == downMatch[0] && upMatch[1] == downMatch[1] {
				xMasCount++
				break
			}
		}
	}

	fmt.Printf("Occurrances of x-Mas: %d\n", xMasCount)
}

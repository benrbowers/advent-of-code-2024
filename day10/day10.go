package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Direction int

const (
	DIR_RIGHT Direction = iota
	DIR_DOWN
	DIR_LEFT
	DIR_UP
	DIR_MAX
)

/*
- Topographic map of digits (0-9)
- Count trailheads and their scores
- Trailheads always start at 0
- Trails move gradually +1 uphill
- Trail heads score = number of 9's that can be reached from it
- Trails can branch, so the solution is at least partially recursive
*/

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:", err)
		panic(err)
	}
	defer inputFile.Close()

	var lines [][]int
	var trailheads [][2]int // [row, col]

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for col, char := range line {
			num := int(char - '0')
			if num == 0 {
				trailheads = append(trailheads, [2]int{len(lines), col})
			}
			nums = append(nums, num)
		}
		lines = append(lines, nums)
	}

	var totalScore int = 0
	var totalRating int = 0

	for _, trailhead := range trailheads {
		var uniquePeaks [][2]int = [][2]int{}
		var rating int = 0
		count9s(lines, trailhead, &uniquePeaks, &rating)
		totalScore += len(uniquePeaks)
		totalRating += rating
	}

	fmt.Printf("Total score: %d\n", totalScore)
	fmt.Printf("Total rating: %d\n", totalRating)
}

func count9s(lines [][]int, pos [2]int, uniquePeaks *[][2]int, rating *int) {
	currElevation := lines[pos[0]][pos[1]]

	if currElevation == 9 {
		if !slices.Contains(*uniquePeaks, pos) {
			*uniquePeaks = append(*uniquePeaks, pos)
		}
		*rating += 1

		return
	}

	for direction := range DIR_MAX {
		nextElevation := getNextElevation(lines, pos, direction)

		if nextElevation == currElevation+1 {
			count9s(lines, moveForward(pos, direction), uniquePeaks, rating)
		}
	}
}

func rotateRight(dir Direction) Direction {
	if dir >= DIR_MAX-1 {
		dir = 0
	} else {
		dir++
	}

	return dir
}

func inBounds(pos [2]int, lines [][]int) bool {
	rows, columns := len(lines), len(lines[0])
	row, col := pos[0], pos[1]

	if row < 0 {
		return false
	}
	if row > rows-1 {
		return false
	}
	if col < 0 {
		return false
	}
	if col > columns-1 {
		return false
	}

	return true
}

func moveForward(pos [2]int, dir Direction) [2]int {
	switch dir {
	case DIR_UP:
		pos[0] -= 1
	case DIR_DOWN:
		pos[0] += 1
	case DIR_LEFT:
		pos[1] -= 1
	case DIR_RIGHT:
		pos[1] += 1
	}

	return pos
}

func getNextElevation(lines [][]int, pos [2]int, direction Direction) int {
	nextPos := moveForward(pos, direction)

	if inBounds(nextPos, lines) {
		return lines[nextPos[0]][nextPos[1]]
	}

	return -1
}

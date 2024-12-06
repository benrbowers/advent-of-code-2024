package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Direction int

const (
	DIR_LEFT Direction = iota
	DIR_UP
	DIR_RIGHT
	DIR_DOWN
	DIR_MAX
)

/*
Count the number of DISTINCT positions the guard
will occupy, INCLUDING the starting position.

The guard ("^") will move with the following rules:
 - Move forward until you reach an obstacle ("#")
 - Turn right 90 degrees
 - Repeat

The guard's path is finsihed when they move out
of bounds of the input area.
*/

func inBounds(pos [2]int, rows, columns int) bool {
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
	case DIR_LEFT:
		pos[1] -= 1
	case DIR_RIGHT:
		pos[1] += 1
	case DIR_DOWN:
		pos[0] += 1
	case DIR_UP:
		pos[0] -= 1
	}

	return pos
}

func rotateRight(dir Direction) Direction {
	if dir >= DIR_MAX-1 {
		dir = 0
	} else {
		dir++
	}

	return dir
}

func getNextChar(pos [2]int, input []string, dir Direction) string {
	lineCount := len(input)
	lineLength := len(input[0])
	newPos := moveForward(pos, dir)

	if inBounds(newPos, lineCount, lineLength) {
		return string(input[newPos[0]][newPos[1]])
	}

	return ""
}

func testLoop(curPos, nextPos, firstPos, secondPos string) bool {
	if firstPos == "" || secondPos == "" {
		return false
	}

	if curPos != firstPos || nextPos != secondPos {
		return false
	}

	return true
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	var guardPos [2]int // row, col
	var originalPos [2]int
	var firstPos string
	var secondPos string
	var lines []string = []string{}
	var distinctPositions []string = []string{} // strings in format "row,col"
	var currentDirection Direction = DIR_UP
	var infiniteLoopCount int = 0

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "^") {
			guardIndex := strings.Index(line, "^")
			guardPos = [2]int{len(lines), guardIndex}
		}

		lines = append(lines, line)
	}

	originalPos = guardPos

	lineCount := len(lines)
	lineLength := len(lines[0])

	for obstacleRow, line := range lines {
		fmt.Printf("%.2f%%\n", float64(obstacleRow)/float64(lineCount)*100.0)
		for obstacleCol, char := range strings.Split(line, "") {
			fmt.Println(obstacleCol)
			if char == "#" {
				continue
			}
			if firstPos != "" {
				guardPos = originalPos
				secondPos = ""
				currentDirection = DIR_UP
			}
			for inBounds(guardPos, lineCount, lineLength) {
				curPos := fmt.Sprintf("%d,%d", guardPos[0], guardPos[1])

				if firstPos == "" {
					firstPos = curPos
				} else if secondPos == "" {
					secondPos = curPos
				}

				if lines[obstacleRow][obstacleCol] == '^' { // So still only count default run
					if !slices.Contains(distinctPositions, curPos) {
						distinctPositions = append(distinctPositions, curPos)
					}
				}

				nextChar := getNextChar(guardPos, lines, currentDirection)
				next := moveForward(guardPos, currentDirection)
				isTestObstacle := next[0] == obstacleRow && next[1] == obstacleCol

				if nextChar == "#" || isTestObstacle {
					currentDirection = rotateRight(currentDirection)
				} else {
					guardPos = moveForward(guardPos, currentDirection)
					nextPos := fmt.Sprintf("%d,%d", guardPos[0], guardPos[1])

					if testLoop(curPos, nextPos, firstPos, secondPos) {
						infiniteLoopCount++
						break
					}
				}
			}
		}
	}

	fmt.Printf("Answer: %d\n", len(distinctPositions))
	fmt.Printf("Configurations with loops: %d\n", infiniteLoopCount)
}

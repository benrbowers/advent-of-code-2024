package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

/*
- Input is individual crop plots
- Letters are the crop types
- Letters of the same type that are adjacent
	are grouped into a "region"
- Smaller regions can exist within larger regions,
	allowing for "holes" in the larger region
- Calculate the sum of the price of fencing for all the regions
- Fencing $ for one region = area * perimeter
- Area = number of cells in the region
- Perimeter = number of sides touching other regions (or the border)

- Finding the regions will need to be recursive
- Finding the perimeter will need to be recursive
	- Ideally, we do both in the same function
*/

type Direction int

const (
	DIR_RIGHT Direction = iota
	DIR_DOWN
	DIR_LEFT
	DIR_UP
	DIR_MAX
)

var rowMax int
var colMax int
var lines []string = []string{}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rowMax = len(lines)
	colMax = len(lines[0])

	var regions [][][2]int = [][][2]int{}
	var totalFencingCost int = 0

	for row := range rowMax {
		for col := range colMax {
			plotAlreadyMapped := false
			for _, region := range regions {
				if slices.Contains(region, [2]int{row, col}) {
					plotAlreadyMapped = true
					break
				}
			}

			if plotAlreadyMapped {
				continue
			}

			var newRegion [][2]int = [][2]int{}

			perimeter := findRegionPerimeter([2]int{row, col}, &newRegion)
			regions = append(regions, newRegion)
			totalFencingCost += perimeter * len(newRegion)
		}
	}

	fmt.Printf("Total fencing cost: %d\n", totalFencingCost)
}

func findRegionPerimeter(startPlot [2]int, currentRegion *[][2]int) int {
	if slices.Contains(*currentRegion, startPlot) {
		return 0
	}

	*currentRegion = append(*currentRegion, startPlot)

	var letter byte = getLetter(startPlot)
	var perimeter int = 0

	for dir := range DIR_MAX {
		nextPos := moveForward(startPlot, dir)
		nextLetter := getLetter(nextPos)

		if nextLetter == letter {
			perimeter += findRegionPerimeter(nextPos, currentRegion)
		} else {
			perimeter += 1
		}
	}

	return perimeter
}

func getLetter(pos [2]int) byte {
	if inBounds(pos) {
		return lines[pos[0]][pos[1]]
	}

	return 0
}

func inBounds(plot [2]int) bool {
	if plot[0] < 0 || plot[0] >= rowMax {
		return false
	}
	if plot[1] < 0 || plot[1] >= colMax {
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

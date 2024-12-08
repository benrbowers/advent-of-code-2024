package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func inBounds(coord [2]int, lineCount, lineLength int) bool {
	if coord[0] < 0 || coord[0] >= lineCount {
		return false
	}
	if coord[1] < 0 || coord[1] >= lineLength {
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

	var lineCount int = 0
	var lineLength int
	var antennas map[byte][][2]int = map[byte][][2]int{} // Map of letters -> coords (row,col)
	var antinodes [][2]int = [][2]int{}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		if lineLength == 0 {
			lineLength = len(line)
		}

		for col := range lineLength {
			if line[col] != '.' {
				antennas[line[col]] = append(
					antennas[line[col]], [2]int{lineCount, col},
				)
			}
		}

		lineCount++
	}

	for _, locations := range antennas {
		for i := 0; i < len(locations)-1; i++ {
			a1 := locations[i]
			for j := i + 1; j < len(locations); j++ {
				a2 := locations[j]

				diff := [2]int{a1[0] - a2[0], a1[1] - a2[1]}

				anti1 := [2]int{a1[0] + diff[0], a1[1] + diff[1]}
				anti2 := [2]int{a2[0] - diff[0], a2[1] - diff[1]}

				if inBounds(anti1, lineCount, lineLength) && !slices.Contains(antinodes, anti1) {
					antinodes = append(antinodes, anti1)
				}
				if inBounds(anti2, lineCount, lineLength) && !slices.Contains(antinodes, anti2) {
					antinodes = append(antinodes, anti2)
				}
			}
		}
	}

	fmt.Printf("Unique antinode locations: %d\n", len(antinodes))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
- Input is compressed file information
- It is a continuous string of digits (0-9)
- The digits alternate between file size and free space
- Each file has an ID equal to its order, starting at 0
- Uncompress into individual file-blocks and free-blocks
- Starting from the left, shift all blocks to the left until there is no free space
- Calculate "checksum": sum of all block IDs * their position
*/

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	input := scanner.Text()

	var memBlocks []int = []int{}

	for i, char := range strings.Split(input, "") {
		size, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf(`"%s" is not a valid size\n`, char)
			panic(err)
		}
		if i%2 == 0 {
			// File block
			for range size {
				memBlocks = append(memBlocks, i/2)
			}
		} else {
			// Free block
			for range size {
				memBlocks = append(memBlocks, -1)
			}
		}
	}

	for slices.Contains(memBlocks, -1) {
		last := memBlocks[len(memBlocks)-1]
		if last != -1 {
			for i, block := range memBlocks {
				if block == -1 {
					memBlocks[i] = last
					break
				}
			}
		}

		memBlocks = memBlocks[:len(memBlocks)-1]
	}

	var checksum int = 0
	for i, block := range memBlocks {
		checksum += block * i
	}

	fmt.Printf("Checksum: %d\n", checksum)
}

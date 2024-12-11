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
	var fileSizes []int = []int{}
	var freeSizes []int = []int{}

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
			fileSizes = append(fileSizes, size)
		} else {
			// Free block
			for range size {
				memBlocks = append(memBlocks, -1)
			}
			freeSizes = append(freeSizes, size)
		}
	}

	// for slices.Contains(memBlocks, -1) {
	// 	last := memBlocks[len(memBlocks)-1]
	// 	if last != -1 {
	// 		for i, block := range memBlocks {
	// 			if block == -1 {
	// 				memBlocks[i] = last
	// 				break
	// 			}
	// 		}
	// 	}

	// 	memBlocks = memBlocks[:len(memBlocks)-1]
	// }

	// var checksum1 int = 0
	// for i, block := range memBlocks {
	// 	checksum1 += block * i
	// }

	// fmt.Printf("Checksum part 1: %d\n", checksum1)

	fmt.Printf("len(fileSizes): %d\n", len(fileSizes))

	for id := len(fileSizes) - 1; id >= 1; id-- {
		fileSize := fileSizes[id]
		if fileSize == 0 {
			continue
		}

		for slot := 0; slot < id-1; slot++ {
			freeSize := freeSizes[slot]
			if fileSize <= freeSize {
				freeSizes[slot] -= fileSize

				fileIndex := slices.Index(memBlocks, id)
				for i := range fileSize {
					memBlocks[fileIndex+i] = -1
				}

				freeIndex := slices.Index(memBlocks, slot+1) - freeSize
				for i := range fileSize {
					memBlocks[freeIndex+i] = id
				}

				break
			}
		}
	}

	var checksum2 int = 0
	for i, block := range memBlocks {
		if block != -1 {
			checksum2 += block * i
		}
	}

	fmt.Printf("Checksum part 2: %d\n", checksum2)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Will probably need a recursive solution for the 75 count
// Something like fn(num) -> fn(num1/2) + fn(num2/2)
// Base case is blinkCount = 75

// Recursive still not fast enough
// Try solving for roots and store in a map: root, blinkCount -> num stones
// If root is already in map, return the value
// If root is not in map, solve for it and store in map

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, " ")

	var rootMap map[string]int = map[string]int{} // [root,blinkCount] -> num stones

	var totalStones int = 0

	for _, num := range nums {
		totalStones += countStones(num, 75, rootMap)
	}

	fmt.Printf("Number of stones: %d\n", totalStones)
}

func countStones(num string, blinkCount int, rootMap map[string]int) int {
	if blinkCount <= 0 {
		return 1
	}

	rootKey := fmt.Sprintf("%s,%d", num, blinkCount)

	if val, exists := rootMap[rootKey]; exists {
		// fmt.Printf("Found %d stones for %s\n", val, rootKey)
		return val
	}

	if len(num)%2 != 0 {
		var stoneCount int
		if num == "0" {
			stoneCount = countStones("1", blinkCount-1, rootMap)
		} else {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Failed to convert num to int:")
				panic(err)
			}
			stoneCount = countStones(strconv.Itoa(numInt*2024), blinkCount-1, rootMap)
		}

		rootMap[rootKey] = stoneCount
		return stoneCount
	}

	half := len(num) / 2
	num1 := num[:half]
	num2 := num[half:]

	if num2[0] == '0' {
		num2int, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Println("Failed to convert num2 to int:")
			panic(err)
		}

		num2 = strconv.Itoa(num2int)
	}

	var leftCount int = 0
	var rightCount int = 0

	leftCount = countStones(num1, blinkCount-1, rootMap)
	rightCount = countStones(num2, blinkCount-1, rootMap)

	totalCount := leftCount + rightCount
	rootMap[rootKey] = totalCount

	return totalCount
}

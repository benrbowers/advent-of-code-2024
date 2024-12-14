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

	var totalStones int = 0

	for _, num := range nums {
		var stoneCount int = 1
		countStones(num, 75, &stoneCount)
		totalStones += stoneCount
	}

	fmt.Printf("Number of stones: %d\n", totalStones)
}

func countStones(num string, blinkCount int, stoneCount *int) {
	if blinkCount <= 0 {
		return
	}

	for len(num)%2 != 0 && blinkCount > 0 {
		if num == "0" {
			num = "1"
		} else {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Failed to convert num to int:")
				panic(err)
			}
			num = strconv.Itoa(numInt * 2024)
		}

		blinkCount -= 1
	}

	if blinkCount <= 0 {
		return
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

	*stoneCount += 1
	countStones(num1, blinkCount-1, stoneCount)
	countStones(num2, blinkCount-1, stoneCount)
}

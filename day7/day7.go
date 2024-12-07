package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The operators "+" and "*" have been stolen from some equations.

Find if any combination of adding and/or multiplying
the the equation numbers can result in the test numbers.

Equate left-to-right, ignoring order of operations.

Input in format:
"test: val1 val2 val3..."
*/

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	var answer int = 0

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")

		testNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf(`"%s" is not a valid test number:`, parts[0])
			panic(err)
		}

		var equationNums []int = []int{}
		for _, numString := range strings.Split(parts[1], " ") {
			num, err := strconv.Atoi(numString)
			if err != nil {
				fmt.Printf(`"%s" is not a valid equation number:`, numString)
				panic(err)
			}
			equationNums = append(equationNums, num)
		}

		n := len(equationNums) - 1 // number of operators

		// Use binary representation, "1010" -> "*+*+"
		// Loop through all 2^n combinations
		for i := 0; i < (1 << n); i++ {
			result := equationNums[0]

			for j := 0; j < n; j++ {
				if (i & (1 << j)) != 0 {
					// The j-th bit is set
					result *= equationNums[j+1]
				} else {
					result += equationNums[j+1]
				}
			}

			if result == testNum {
				answer += testNum
				break
			}
		}
	}

	fmt.Printf("Total of valid test nums: %d\n", answer)
}

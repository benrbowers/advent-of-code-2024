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

		var equationNums []int
		for _, numString := range strings.Split(parts[1], " ") {
			num, err := strconv.Atoi(numString)
			if err != nil {
				fmt.Printf(`"%s" is not a valid equation number:`, numString)
				panic(err)
			}
			equationNums = append(equationNums, num)
		}

		var validCombo bool = false
		n := len(equationNums) - 1 // number of operators

		for x := 0; x < (1 << n); x++ {
			if validCombo {
				break
			}
			var concatNums []int
			var concatActive bool = false

			for y := 0; y < n; y++ {
				if (x & (1 << y)) != 0 {
					// The y-th bit is set, so concat
					var num1 string
					if concatActive {
						num1 = strconv.Itoa(concatNums[len(concatNums)-1])
					} else {
						num1 = strconv.Itoa(equationNums[y])
					}
					num2 := strconv.Itoa(equationNums[y+1])
					newNum, err := strconv.Atoi(num1 + num2)
					if err != nil {
						fmt.Printf(`"%s" is not a number:`, num1+num2)
						panic(err)
					}

					if concatActive {
						concatNums[len(concatNums)-1] = newNum
					} else {
						concatNums = append(concatNums, newNum)
					}

					concatActive = true
				} else {
					concatActive = false
					concatNums = append(concatNums, equationNums[y])
				}
			}

			// Use binary representation, "1010" -> "*+*+"
			// Loop through all 2^n combinations
			nc := len(concatNums) - 1

			for i := 0; i < (1 << nc); i++ {
				result := concatNums[0]

				for j := 0; j < nc; j++ {
					if (i & (1 << j)) != 0 {
						// The j-th bit is set
						// fmt.Print("*")
						result *= concatNums[j+1]
					} else {
						// fmt.Print("+")
						result += concatNums[j+1]
					}
				}

				// fmt.Println()
				if result == testNum {
					answer += testNum
					validCombo = true
					break
				}
			}

			// bits := strings.Split(fmt.Sprintf("%b\n", x), "")
			// slices.Reverse(bits)
			// fmt.Println(strings.Join(bits, ""))
			// fmt.Println(concatNums)
		}
	}

	fmt.Printf("Total of valid test nums: %d\n", answer)
}

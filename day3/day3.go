package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
mul func expects 2, 1-3 digit numbers and multiplies them
ignore input with whitespace or any non-number characters
*/

func mul(x, y int) int {
	if x > 999 {
		panic("X has too many digits: " + strconv.Itoa(x))
	}
	if y > 999 {
		panic("Y has too many digits: " + strconv.Itoa(y))
	}

	return x * y
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt: ")
		panic(err)
	}
	defer inputFile.Close()

	mulPattern := regexp.MustCompile(
		"(mul\\((0|[1-9][0-9]{0,2}),(0|[1-9][0-9]{0,2})\\))|(do\\(\\))|(don't\\(\\))",
	)

	var answer int = 0
	var mulEnabled bool = true

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		matches := mulPattern.FindAllString(line, -1)

		for _, match := range matches {
			if match == "do()" {
				mulEnabled = true
			} else if match == "don't()" {
				mulEnabled = false
			} else if mulEnabled {
				match = match[4 : len(match)-1]
				nums := strings.Split(match, ",")

				x, err := strconv.Atoi(nums[0])
				if err != nil {
					fmt.Println("X is not a number:")
					panic(err)
				}

				y, err := strconv.Atoi(nums[1])
				if err != nil {
					fmt.Println("Y is not a number:")
					panic(err)
				}

				answer += mul(x, y)
			}
		}
	}

	fmt.Println("Sum of muls: " + strconv.Itoa(answer))
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not open input:")
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var safeCount int = 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		var levels []int = []int{}

		for _, digit := range nums {
			num, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println(digit + " is not a valid integer:")
				panic(err)
			}
			levels = append(levels, num)
		}

		var increasing bool = false
		var faultCount int = 0

		for i, num := range levels {
			if i == 0 && num < levels[1] {
				increasing = true
			}
			if faultCount > 1 {
				break
			}

			if i < len(levels)-1 {
				if increasing {
					if num > levels[i+1] {
						faultCount++
						continue
					}
				} else {
					if num < levels[i+1] {
						faultCount++
						continue
					}
				}

				if num == levels[i+1] {
					faultCount++
					continue
				}

				if math.Abs(float64(levels[i+1]-num)) > 3 {
					faultCount++
					continue
				}
			} else {
				safeCount++
			}
		}
	}

	fmt.Println("Safe count: " + strconv.Itoa(safeCount))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	for range 25 {
		var newNums []string = []string{}
		for _, num := range nums {
			if num == "0" {
				newNums = append(newNums, "1")
			} else if len(num)%2 == 0 {
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

				// fmt.Printf("num: %s\n", num)
				// fmt.Printf("num1: %s, num2: %s\n", num1, num2)
				newNums = append(newNums, num1, num2)
			} else {
				numInt, err := strconv.Atoi(num)
				if err != nil {
					fmt.Println("Failed to convert num to int:")
					panic(err)
				}
				newNums = append(newNums, strconv.Itoa(numInt*2024))
			}
		}
		nums = newNums
	}

	fmt.Printf("Number of stones: %d\n", len(nums))
}

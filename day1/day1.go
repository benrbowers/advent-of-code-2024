package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	var distanceScore int = 0
	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		list2 = append(list2, num2)
	}

	if len(list1) != len(list2) {
		panic("Lists are not of equal length")
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		distanceScore += int(math.Abs(float64(list1[i] - list2[i])))
	}

	var similarityScore int = 0
	var rightCounts map[int]int = map[int]int{}

	for _, num := range list2 {
		rightCounts[num]++
	}

	for _, num := range list1 {
		similarityScore += num * rightCounts[num]
	}

	fmt.Println("Distance Score:", distanceScore)
	fmt.Println("Similarity Score:", similarityScore)
}

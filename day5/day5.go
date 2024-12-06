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
First section to parse is rules regarding page number order
in the form "X|Y", where X must come BEFORE Y,
regardless of actual numeric value

Second section is list of page sequences to be checked
against rules from the first section

Sections are separated by new line
*/

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	var rulesList [][]string = [][]string{}
	var validSequences [][]string = [][]string{}
	var isFirstSection bool = true

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		if isFirstSection {
			if line == "" {
				isFirstSection = false
				continue
			}

			ruleNums := strings.Split(line, "|")

			if len(ruleNums) != 2 {
				panic(fmt.Sprintf(`Invalid rule format: "%s" \n`, line))
			}

			rulesList = append(rulesList, ruleNums)
		} else {
			pageNums := strings.Split(line, ",")
			var sequenceValid bool = true

			for _, rule := range rulesList {
				if slices.Contains(pageNums, rule[0]) && slices.Contains(pageNums, rule[1]) {
					if slices.Index(pageNums, rule[0]) > slices.Index(pageNums, rule[1]) {
						sequenceValid = false
						break
					}
				}
			}

			if sequenceValid {
				validSequences = append(validSequences, pageNums)
			}
		}
	}

	// Sum of middle page numbers
	var answer int = 0

	for _, sequence := range validSequences {
		middle := sequence[len(sequence)/2]
		middleNum, err := strconv.Atoi(middle)
		if err != nil {
			panic(fmt.Sprintf(`Invalid page number: "%s" \n`, middle))
		}

		answer += middleNum
	}

	fmt.Printf("Answer: %d\n", answer)
	fmt.Println(validSequences)
}

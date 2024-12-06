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

func swapValues(slice []string, val1, val2 string) {
	index1 := slices.Index(slice, val1)
	index2 := slices.Index(slice, val2)

	if index1 == -1 || index2 == -1 {
		return
	}

	slice[index1], slice[index2] = slice[index2], slice[index1]
}

func isValidSequence(sequence []string, rulesList [][]string) bool {
	for _, rule := range rulesList {
		if slices.Contains(sequence, rule[0]) && slices.Contains(sequence, rule[1]) {
			if slices.Index(sequence, rule[0]) > slices.Index(sequence, rule[1]) {
				return false
			}
		}
	}

	return true
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input.txt:")
		panic(err)
	}
	defer inputFile.Close()

	var rulesList [][]string = [][]string{}
	var validSequences [][]string = [][]string{}
	var invalidSequences [][]string = [][]string{}
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
			if isValidSequence(pageNums, rulesList) {
				validSequences = append(validSequences, pageNums)
			} else {
				invalidSequences = append(invalidSequences, pageNums)
			}
		}
	}

	// Sum of middle page numbers
	var answer int = 0
	var invalidAnswer int = 0

	for _, sequence := range validSequences {
		middle := sequence[len(sequence)/2]
		middleNum, err := strconv.Atoi(middle)
		if err != nil {
			panic(fmt.Sprintf(`Invalid page number: "%s" \n`, middle))
		}

		answer += middleNum
	}

	for _, sequence := range invalidSequences {
		for !isValidSequence(sequence, rulesList) {
			// While sequence is not valid, swap values according to rules
			for _, rule := range rulesList {
				if slices.Contains(sequence, rule[0]) && slices.Contains(sequence, rule[1]) {
					if slices.Index(sequence, rule[0]) > slices.Index(sequence, rule[1]) {
						swapValues(sequence, rule[0], rule[1])
					}
				}
			}
		}
	}

	for _, sequence := range invalidSequences {
		middle := sequence[len(sequence)/2]
		middleNum, err := strconv.Atoi(middle)
		if err != nil {
			panic(fmt.Sprintf(`Invalid page number: "%s" \n`, middle))
		}

		invalidAnswer += middleNum
	}

	fmt.Printf("Answer: %d\n", answer)
	fmt.Printf("Invalid answer: %d\n", invalidAnswer)
}

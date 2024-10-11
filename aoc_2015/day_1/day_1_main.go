package main

import (
	"fmt"
	"os"
)

func main() {

	fileData, err := os.ReadFile("aoc_2015/day_1/day1data.txt")
	if err != nil {
		return
	}

	constant := 0
	result := 0
	charPosition := 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '(' {
			result++
		} else if fileData[i] == ')' {
			result--
		}

		if constant == 0 && result == -1 {
			constant = 1
			charPosition = i + 1
		}
	}

	fmt.Println("Result is: ", result)
	fmt.Println("Position is: ", charPosition)
}

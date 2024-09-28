package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, err := os.Open("aoc_2023/day_1/day1data.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstNumber(line)*10 + lastNumber(line)
	}

	fmt.Printf("The sum is: %v", sum)
}

func firstNumber(s string) int {
	str := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		str += string(s[i])

		for i, d := range digits {
			if strings.HasSuffix(str, d) {
				return i + 1
			}
		}

	}
	log.Fatal("first digit not found")
	return -1
}

func lastNumber(s string) int {
	str := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		str = string(s[i]) + str

		for i, d := range digits {
			if strings.HasPrefix(str, d) {
				return i + 1
			}
		}
	}
	log.Fatal("last digit not found")
	return -1
}

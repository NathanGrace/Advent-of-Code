package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ColorExtractions struct {
	Red, Blue, Green int
}

func main() {
	file, err := os.Open("aoc_2023/day_2/day2testingdata.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	goodGameIdsSum := 0
	isPossible := false
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line)
	}

	fmt.Printf("this game is possible: %v\n", isPossible)
	fmt.Println(goodGameIdsSum)
}

func parseLine(line string) (int, []ColorExtractions) {
	parseGame := strings.Split(line, ":")
	var game int
	fmt.Sscanf(parseGame[0], "Game %d", &game)

	colorExtractions := []ColorExtractions{}

	//for colors := range strings.Split(parseGame[1], ";") {
	//}

	return game, colorExtractions
}

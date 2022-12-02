package main

import (
	"fmt"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day02go1/input.txt")
	fmt.Println(rpsScore(input))
}

func rpsScore(ss []string) int {
	var score int
	for _, s := range ss {
		score += scorePair(s)
	}
	return score
}

func scorePair(s string) int {
	return scoreMap[s]
}

var scoreMap = map[string]int{
	"A X": 4, // 1 + 3
	"A Y": 8, // 2 + 6
	"A Z": 3, // 3 + 0
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 7, // 1 + 6
	"C Y": 2, // 2 + 0
	"C Z": 6, // 3 + 3
}

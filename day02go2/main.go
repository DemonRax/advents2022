package main

import (
	"fmt"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day02go2/input.txt")
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
	"A X": 3, // 0 + 3
	"A Y": 4, // 3 + 1
	"A Z": 8, // 6 + 2
	"B X": 1, // 0 + 1
	"B Y": 5, // 3 + 2
	"B Z": 9, // 6 + 3
	"C X": 2, // 0 + 2
	"C Y": 6, // 3 + 3
	"C Z": 7, // 6 + 1
}

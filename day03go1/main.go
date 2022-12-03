package main

import (
	"fmt"

	"advents2022/util"
)

var priorities map[rune]int

func init() {
	priorities = make(map[rune]int, 52)
	for i, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		priorities[c] = i + 1
	}
}

func main() {
	input := util.ReadFile("./day03go1/input.txt")
	fmt.Println(prioSum(input))
}

func prioSum(ss []string) int {
	var score int
	for _, s := range ss {
		middle := len(s) / 2
		c := common(s[:middle], s[middle:])
		score += priorities[c]
	}
	return score
}

func common(s1, s2 string) rune {
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				return c1
			}
		}
	}
	return 0
}

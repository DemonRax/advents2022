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
	input := util.ReadFile("./day03go2/input.txt")
	fmt.Println(badgesSum(input))
}

func badgesSum(ss []string) int {
	var count, score int
	set := make(map[rune]int, 52)
	for _, s := range ss {
		rucksack := make(map[rune]int, 52)
		for _, c := range s {
			rucksack[c] = 1
		}
		for c := range rucksack {
			set[c] = set[c] + 1
		}
		count++
		if count == 3 {
			count = 0
			for c, i := range set {
				if i == 3 {
					score += priorities[c]
					set = make(map[rune]int, 52)
					break
				}
			}
		}
	}
	return score
}

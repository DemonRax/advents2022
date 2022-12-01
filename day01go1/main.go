package main

import (
	"fmt"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day01go1/input.txt")
	fmt.Println(maxCalories(input))
}

func maxCalories(ss []string) int {
	var sum, max int
	for _, s := range ss {
		if s == "" {
			sum = 0
		}
		sum += util.ToInt(s)
		if sum > max {
			max = sum
		}
	}
	return max
}

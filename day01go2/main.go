package main

import (
	"fmt"
	"sort"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day01go2/input.txt")
	input = append(input, "")
	fmt.Println(topThree(input))
}

func topThree(ss []string) int {
	calories := make([]int, 10)
	var sum int
	for _, s := range ss {
		sum += util.ToInt(s)
		if s == "" {
			calories = append(calories, sum)
			sum = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	for i := 0; i < 3; i++ {
		sum += calories[i]
	}
	return sum
}

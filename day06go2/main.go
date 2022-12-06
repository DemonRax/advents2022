package main

import (
	"advents2022/util"
	"fmt"
)

func main() {
	input := util.ReadFile("./day06go2/input.txt")
	fmt.Println(firstUnique(input[0], 14))
}

func firstUnique(s string, size int) int {
	for i := range s {
		if i < size {
			continue
		}
		if unique(s[i-size : i]) {
			return i
		}
	}
	return 0
}

func unique(s string) bool {
	set := make(map[rune]interface{}, len(s))
	for _, c := range s {
		set[c] = struct{}{}
	}
	return len(set) == len(s)
}

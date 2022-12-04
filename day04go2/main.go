package main

import (
	"fmt"
	"strings"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day04go2/input.txt")
	fmt.Println(overlaps(input))
}

func overlaps(ss []string) int {
	var count int
	for _, s := range ss {
		if doesOverlap(s) {
			count++
		}
	}
	return count
}

func doesOverlap(s string) bool {
	a, b, x, y := convert(s)
	ii := util.Interval{a, b}
	jj := util.Interval{x, y}
	return !ii.After(jj) && !ii.Before(jj)
}

func convert(s string) (int, int, int, int) {
	ss := strings.Split(s, ",")
	i1 := strings.Split(ss[0], "-")
	a, b := util.ToInt(i1[0]), util.ToInt(i1[1])
	i2 := strings.Split(ss[1], "-")
	x, y := util.ToInt(i2[0]), util.ToInt(i2[1])
	return a, b, x, y
}

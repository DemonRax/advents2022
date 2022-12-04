package main

import (
	"fmt"
	"strings"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day04go1/input.txt")
	fmt.Println(fully(input))
}

func fully(ss []string) int {
	var count int
	for _, s := range ss {
		if stringInclude(s) {
			count++
		}
	}
	return count
}

func stringInclude(s string) bool {
	ss := strings.Split(s, ",")
	i1 := strings.Split(ss[0], "-")
	a, b := util.ToInt(i1[0]), util.ToInt(i1[1])
	i2 := strings.Split(ss[1], "-")
	x, y := util.ToInt(i2[0]), util.ToInt(i2[1])
	return intervalsInclude(a, b, x, y)
}

func intervalsInclude(a, b, x, y int) bool {
	return intervalInclude(a, b, x, y) || intervalInclude(x, y, a, b)
}

func intervalInclude(a, b, x, y int) bool {
	return x >= a && y <= b
}

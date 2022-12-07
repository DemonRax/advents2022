package main

import (
	"advents2022/util"
	"fmt"
)

func main() {
	input := util.ReadFile("./day07go1/input.txt")
	fmt.Println(dirSize(input))
}

func dirSize(ss []string) int {
	ip := util.NewInterpreter()
	for _, s := range ss {
		ip.ScanLine(s)
	}
	var sum int
	for _, size := range ip.DirSizes() {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

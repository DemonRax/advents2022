package main

import (
	"advents2022/util"
	"fmt"
)

func main() {
	input := util.ReadFile("./day07go2/input.txt")
	fmt.Println(dirSize(input))
}

func dirSize(ss []string) int {
	ip := util.NewInterpreter()
	for _, s := range ss {
		ip.ScanLine(s)
	}

	free := 70000000 - ip.TotalSize()
	min := 30000000
	for _, size := range ip.DirSizes() {
		if size+free > 30000000 && size < min {
			min = size
		}
	}
	return min
}

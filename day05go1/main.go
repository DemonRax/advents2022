package main

import (
	"fmt"
	"strings"

	"advents2022/util"
)

func main() {
	input := util.ReadFile("./day05go1/input.txt")
	fmt.Println(top(input))
}

func top(ss []string) string {
	var res strings.Builder
	for _, stack := range operate(read(ss))[1:] {
		res.Write(stack[:1])
	}
	return res.String()
}

func operate(stacks [][]byte, ops [][]int) [][]byte {
	for _, op := range ops {
		qty, from, to := op[0], op[1], op[2]
		cut := make([]byte, qty)
		for i := 0; i < qty; i++ {
			cut[i] = stacks[from][qty-i-1]
		}
		stacks[to] = append(cut, stacks[to]...)
		stacks[from] = stacks[from][qty:]
	}
	return stacks
}

func read(ss []string) ([][]byte, [][]int) {
	stacks := make([]string, 0, len(ss))
	ops := make([]string, 0, len(ss))

	target := &stacks

	for _, ss := range ss {
		if ss == "" {
			target = &ops
			continue
		}
		*target = append(*target, ss)
	}

	return convertStacks(stacks), convertOps(ops)
}

func convertStacks(ss []string) [][]byte {
	size := len(strings.ReplaceAll(ss[util.Last(ss)], " ", ""))
	stacks := make([][]byte, size+1)

	//for line := util.Last(ss) - 1; line >= 0; line-- {
	for line := 0; line < util.Last(ss); line++ {
		for i := 0; i < size; i++ {
			c := ss[line][i*4+1]
			if c != 32 {
				stacks[i+1] = append(stacks[i+1], c)
			}
		}
	}
	return stacks
}

func convertOps(ss []string) [][]int {
	ops := make([][]int, len(ss))
	for i, s := range ss {
		cuts := strings.Split(s, " ")
		ops[i] = []int{util.ToInt(cuts[1]), util.ToInt(cuts[3]), util.ToInt(cuts[5])}

	}
	return ops
}

package main

import (
	"advents2022/util"
	"fmt"
)

func main() {
	input := util.ReadFile("./day08go1/input.txt")
	fmt.Println(visibleTrees(input))
}

func visibleTrees(ss []string) int {
	var sum int
	grid := generateGrid(ss)
	for i, row := range grid {
		for j := range row {
			if visibleRow(i, j, grid) || visibleCol(i, j, grid) {
				sum++
			}
		}
	}
	return sum
}

func generateGrid(ss []string) [][]int {
	grid := make([][]int, len(ss))
	for i, s := range ss {
		grid[i] = make([]int, len(s))
		for j, c := range s {
			grid[i][j] = util.ToInt(string(c))
		}
	}
	return grid
}

func visibleRow(i, j int, grid [][]int) bool {
	if j == 0 || j == len(grid[0])-1 {
		return true
	}
	tree := grid[i][j]
	visible := true
	for x := range grid[i] {
		if x == j {
			if visible {
				return visible
			}
			visible = true
			continue
		}
		if tree <= grid[i][x] {
			visible = false
		}
	}
	return visible
}

func visibleCol(i, j int, grid [][]int) bool {
	if i == 0 || i == len(grid)-1 {
		return true
	}
	tree := grid[i][j]
	visible := true
	for x := 0; x < len(grid); x++ {
		if x == i {
			if visible {
				return visible
			}
			visible = true
			continue
		}
		if tree <= grid[x][j] {
			visible = false
		}
	}
	return visible
}

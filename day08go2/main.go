package main

import (
	"advents2022/util"
	"fmt"
)

func main() {
	input := util.ReadFile("./day08go2/input.txt")
	fmt.Println(highestScore(input))
}

func highestScore(ss []string) int {
	var high int
	grid := generateGrid(ss)
	for i, row := range grid {
		for j := range row {
			if score := scoreLeft(i, j, grid) * scoreRight(i, j, grid) * scoreUp(i, j, grid) * scoreDown(i, j, grid); score > high {
				high = score
			}
		}
	}
	return high
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

func scoreUp(i, j int, grid [][]int) int {
	if i == 0 {
		return 0
	}
	return score(grid[i][j], i-1, 0, j, j, grid)
}

func scoreDown(i, j int, grid [][]int) int {
	if i == len(grid)-1 {
		return 0
	}
	return score(grid[i][j], i+1, len(grid)-1, j, j, grid)
}

func scoreRight(i, j int, grid [][]int) int {
	if j == len(grid[0])-1 {
		return 0
	}
	return score(grid[i][j], i, i, j+1, len(grid[0])-1, grid)
}

func scoreLeft(i, j int, grid [][]int) int {
	if j == 0 {
		return 0
	}
	return score(grid[i][j], i, i, j-1, 0, grid)
}

var funcLess = func(i, max int) bool {
	return i <= max
}

var funcMore = func(i, max int) bool {
	return i >= max
}

func score(tree, i0, iN, j0, jN int, grid [][]int) int {
	iStep := 1
	iPred := funcLess
	if iN < i0 {
		iStep = -1
		iPred = funcMore
	}
	jStep := 1
	jPred := funcLess
	if jN < j0 {
		jStep = -1
		jPred = funcMore
	}

	var score int
	for i := i0; iPred(i, iN); i += iStep {
		for j := j0; jPred(j, jN); j += jStep {
			score++
			if tree <= grid[i][j] {
				return score
			}
		}
	}
	return score
}

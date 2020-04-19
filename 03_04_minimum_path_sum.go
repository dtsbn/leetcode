package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	matrix := make([][]int, m)

	// fill first row
	for sum, y := 0, 0; y < m; y++ {
		matrix[y] = make([]int, n)
		sum += grid[y][0]
		matrix[y][0] = sum
	}

	//fill first column
	for sum, x := 0, 0; x < n; x++ {
		sum += grid[0][x]
		matrix[0][x] = sum
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			matrix[y][x] = min(matrix[y][x-1], matrix[y-1][x]) + grid[y][x]
		}
	}

	return matrix[m-1][n-1]
}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	input2 := [][]int{
		{3, 8, 6, 0, 5, 9, 9, 6, 3, 4, 0, 5, 7, 3, 9, 3},
		{0, 9, 2, 5, 5, 4, 9, 1, 4, 6, 9, 5, 6, 7, 3, 2},
		{8, 2, 2, 3, 3, 3, 1, 6, 9, 1, 1, 6, 6, 2, 1, 9},
	}

	fmt.Println(minPathSum(input))
	fmt.Println(minPathSum(input2))
}

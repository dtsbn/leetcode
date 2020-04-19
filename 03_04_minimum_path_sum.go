package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(grid [][]int, i int, j int, path int) int {
	if i < len(grid)-1 && j < len(grid[0])-1 {
		return grid[i][j] + min(dfs(grid, i+1, j, path), dfs(grid, i, j+1, path))
	} else if i < len(grid)-1 {
		return grid[i][j] + dfs(grid, i+1, j, path)
	} else if j < len(grid[0])-1 {
		return grid[i][j] + dfs(grid, i, j+1, path)
	} else {
		return grid[i][j]
	}
}

func minPathSum(grid [][]int) int {
	return dfs(grid, 0, 0, 0)
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

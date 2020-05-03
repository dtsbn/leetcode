package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{17, 12, 5, -6, 12, 4, 17, -5, 2, -3, 2, 4, 5, 16, -3, -4, 15, 15, -4, -5, -6},
	}

	for _, i := range input {
		fmt.Println(singleNumber(i))
	}

}

func singleNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) > 1 && nums[0] != nums[1] {
		return nums[0]
	}

	if len(nums) > 1 && nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return -1
}

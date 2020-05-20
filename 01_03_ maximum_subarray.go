package main

import "fmt"

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(input))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxOfThree(a, b, c int) int {
	return max(max(a, b), c)
}

func maxSubArray(nums []int) int {
	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxSubArraySum(arr []int, l int, h int) int {
	if l == h {
		return arr[l]
	}

	middle := (l + h) / 2

	return maxOfThree(maxSubArraySum(arr, l, middle),
		maxSubArraySum(arr, middle+1, h),
		maxCrossingSum(arr, l, middle, h))
}

func maxCrossingSum(arr []int, l int, m int, h int) int {

}

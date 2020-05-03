package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	input := 2

	fmt.Println(isHappy(input))

}

func isHappy(n int) bool {
	fmt.Println("Input: ", n)

	if n == 0 || n == 2 || n == 4 || n == 6 {
		return false
	} else if n == 1 {
		return true
	}

	str := strconv.Itoa(n)

	nums := make([]int, utf8.RuneCountInString(str))
	for pos, c := range str {
		nums[pos], _ = strconv.Atoi(string(c))
	}

	sum := 0
	for _, n := range nums {
		sum += n * n
	}

	return isHappy(sum)
}

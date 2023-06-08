package main

import (
	"fmt"
	"sort"
)

// 暴力法
func sortedSquares(nums []int) []int {
	if nums == nil {
		return nil
	}
	newnums := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		newnums = append(newnums, nums[i]*nums[i])
	}
	sort.Ints(newnums)
	return newnums
}

// 双指针法
func sortedSquares2(nums []int) []int {
	if nums == nil {
		return nil
	}
	newnums := make([]int, len(nums))
	fast := len(nums) - 1
	slow := 0

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[fast]*nums[fast] > nums[slow]*nums[slow] {
			newnums[i] = nums[fast] * nums[fast]
			fast--
		} else {
			newnums[i] = nums[slow] * nums[slow]
			slow++
		}
	}
	return newnums
}
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	res := sortedSquares2(nums)
	fmt.Println(res)
}

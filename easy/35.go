package main

import (
	"fmt"
)

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法
//输入: nums = [1,3,5,6], target = 5
//输出: 2

// 先排序，再找字母是否存在
// 存在，返回位置
// 不存在，插入并返回位置
func searchInsert(nums []int, target int) int {
	//	var b bool
	//	fmt.Println(nums)

	for i, n := range nums {
		switch true {
		case target > n && len(nums) == i+1:
			nums = append(nums[:i], target)
			return i + 1
		case target == n:
			return i
		case target < n:

			return i
		default:
			continue
		}

	}
	fmt.Println(nums)
	return 0
}
func main() {
	targent := 7
	nums := []int{1, 3, 5, 6}
	n := searchInsert(nums, targent)
	fmt.Println(n)
}

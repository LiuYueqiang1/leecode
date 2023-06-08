package main

import "fmt"

// 左闭右闭区间
func errsearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := left + (right-left)/2
	for nums[middle] != target {

		if nums[middle] > target {
			right = middle - 1
			middle = left + (right-left)/2
		} else if nums[middle] < target {
			left = middle + 1
			middle = left + (right-left)/2
		}
		return middle
	}
	return -1
}
func search(nums []int, target int) int {
	//左闭右闭
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

// 左闭右开
func search2(nums []int, target int) int {
	//左闭右开
	left := 0
	right := len(nums) - 1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle
		} else {
			return middle
		}
	}
	return -1
}
func main() {
	ret := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println(ret)
}

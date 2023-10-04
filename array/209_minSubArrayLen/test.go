package main

import "fmt"

func minSubArrayLen2(target int, nums []int) int {
	//起始位置、终止位置

	//终止位置++
	//如果终止位置-起始位置的数组中的和 >= target
	// 起始位置++
	//直到 终止位置到达终点
	start, end := 0, 0
	// 记录一下此时的长度
	lens := make([]int, 0)
	for end < len(nums) {
		for sum(nums[start:end+1]) >= target {
			lens = append(lens, len(nums[start:end+1]))
			start++
		}
		end++
	}
	if len(lens) == 0 {
		return 0
	}
	return mins(lens)
}

func sum(nums []int) int {
	numsum := 0
	for _, v := range nums {
		numsum += v
	}
	return numsum
}

func mins(nums []int) int {
	minlen := nums[0]
	for _, v := range nums {
		if minlen >= v {
			minlen = v
		}
	}
	return minlen
}

func main() {
	res := minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(res)
}

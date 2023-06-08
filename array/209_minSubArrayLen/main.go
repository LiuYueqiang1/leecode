package main

import "fmt"

// 给定一个含有 n 个正整数的数组和一个正整数 s ，
// 找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
func minSubArrayLen(s int, nums []int) int {
	minLen := len(nums) + 1 // 最小长度
	left, sum := 0, 0       // 左指针和子数组和
	for right := 0; right < len(nums); right++ {
		sum += nums[right] // 将当前元素加入子数组
		for sum >= s {     // 子数组满足条件
			minLen = min(minLen, right-left+1) // 更新最小长度
			sum -= nums[left]                  // 将左端点移除子数组
			left++                             // 更新左指针
		}
	}
	if minLen == len(nums)+1 { // 没有符合条件的子数组
		return 0
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(res)
}

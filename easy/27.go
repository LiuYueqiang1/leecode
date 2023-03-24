package main

import "fmt"

// 给你一个数组 nums 和一个值 val，你需要 原地
// 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]
func removeElement(nums []int, val int) int {
	//for i1, n1 := range nums {
	//	if n1 == val {
	//		fmt.Println(len(nums))
	//		nums = append(nums[:i1], nums[i1+1:]...)
	//	}
	//}
tag1:
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			goto tag1
		}
	}
	fmt.Println(nums)
	return len(nums)
}
func main() {
	nums := []int{0, 2, 2, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)
}

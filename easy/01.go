package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	lenght := len(nums)
	var a = make([]int, 0)
	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if nums[i]+nums[j] == 6 {
				a = append(a, i, j)
				//a = append(a, j)
				fmt.Println(a)
				return a
			}
		}
	}
	return a
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	twoSum(nums, target)

}

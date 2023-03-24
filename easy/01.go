package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	lenght := len(nums)
	var a = make([]int, 0)
	for i := 0; i < lenght; i++ {
		for j := 1; j < lenght; j++ {
			if nums[i]+nums[j] == target {
				a = append(a, i, j)
				//a = append(a, j)
				fmt.Println(a)
				return a
			}
		}
	}
	return a
	//这个代码有一个缺点，就是2 5 5 1 ，target=10时输出的是1 1，而不是 1 2
}
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if val, exist := m[nums[i]]; exist {
			return []int{val, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}
func main() {
	nums := []int{2, 5, 5, 11}
	target := 10
	//twoSum(nums, target)
	twoSum2(nums, target)

}

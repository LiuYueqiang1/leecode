package main

import "fmt"

//1. 两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 查找类型的首先想到哈希
// 数组，set，map
func twoSum(nums []int, target int) []int {
	//创建一个map 存放值和下标
	m := make(map[int]int, 0)
	for i, v := range nums {
		if subindex, ok := m[target-v]; ok { //查看 target-当前值 是否在map中存在
			return []int{subindex, i} //如果存在，返回 target-当前值，当前值 的下标
		} else {
			//如果不存在的话，将map加到里面去
			m[v] = i
		}
	}
	return nil
}
func main() {
	ret := twoSum([]int{2, 7, 5, 6}, 9)
	fmt.Println(ret)
}

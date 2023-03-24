package main

import "fmt"

// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，
// 使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//输入：nums = [0,0,1,1,1,2,2,3,3,4]
//输出：5, nums = [0,1,2,3,4]
//解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

func removeDuplicates1(nums []int) int {
	//1、遍历得到所有的元素
	//2、将每个元素向后寻找，如果有相同的元素则删除，且更新切片
tag1:
	for j := 0; j < len(nums); j++ {
		for i := j + 1; i < len(nums); i++ {
			//tag1:
			if nums[j] == nums[i] {
				nums = append(nums[:i], nums[i+1:]...)
				goto tag1
			}
		}
	}

	fmt.Println(nums)
	fmt.Println(len(nums))
	return len(nums)
}

// leecode上的
// 思路：
// 给定一个计数器，用计数器做一个新的数组
// 如果下一个跟前一个不相等的话，新的数组就是原来的，则计数器+1
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i] //从第二个数开始写入新数组，第一个不变
			j++
		}
	}
	return j
}
func main() {
	var nums = make([]int, 0, 20)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret := removeDuplicates2(nums)
	fmt.Println(ret)
}

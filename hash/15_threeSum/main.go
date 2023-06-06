package main

import (
	"fmt"
	"sort"
)

// 第15题. 三数之和
// 给你一个包含 n 个整数的数组 nums，
// 判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}

		if i > 0 && n1 == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2 := nums[left]
			n3 := nums[right]
			if n1+n2+n3 == 0 {
				ret = append(ret, []int{n1, n2, n3})
				//去重
				for left < right && n2 == nums[left] {
					//从它自己等于它自己开始，left++后看是否相等，相等再++，直到遇到nums[left]！=n2，第一次相等数为止
					left++
				}
				// 不能用if continue 操作，因为 left从i+1开始索引， 可能会跟上一个i的索引进行比较，这样的话就得另外考虑了。
				// 而 i 则不会，它本身就是跟自己的，不会出现上面的情况
				for left < right && n3 == nums[right] {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	// 找出a + b + c = 0
	// a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < len(nums)-2; i++ {
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重a
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				// 找到之后才去重，否则不用去重
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
func main() {
	ret := threeSum2([]int{-1, 0, 1, 1, 1, 1, 2, -1, -1, -1, -1, -4})
	fmt.Println(ret)
}

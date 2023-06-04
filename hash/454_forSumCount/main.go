package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := make(map[int]int, 0)
	count := 0 //出现的次数
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			sum := -(v1 + v2)
			m1[sum]++
		}
	}
	fmt.Println(m1)
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			sum2 := v3 + v4
			//这样的话遍历出来的每个数都会去count++
			//如果在map中存在两个的话直接加上两个
			//否则map[else]=0，加的也是0
			count += m1[sum2]
			//错误的：
			//for m1[sum2] != 0 {
			//	if _, ok := m1[sum2]; ok {
			//		count++
			//		m1[sum2]--
			//		fmt.Println(m1)
			//	} else {
			//		m1[sum2]--
			//	}
			//
			//}
		}
	}

	return count
}

func main() {
	ret := fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1})
	fmt.Println(ret)
}

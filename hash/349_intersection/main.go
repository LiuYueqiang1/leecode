package main

import "fmt"

/*
349. 两个数组的交集
题意：给定两个数组，编写一个函数来计算它们的交集。
说明：输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。
*/

// 在go语言中使用map来模拟set
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]interface{})
	ret := make([]int, 0)
	for _, v := range nums1 { //拿到 nums1 中所有的值
		if _, ok := set[v]; !ok { //*** 如果不在map中，则加到 map 中，只有key，不需要 value ，所以map[key]=空结构体
			set[v] = struct {
			}{}
		}
	}
	// fmt.Println(set) //map[1:{} 2:{}]
	for _, v := range nums2 {
		if _, ok := set[v]; ok { //*** 如果在map中
			ret = append(ret, v) // 如果 nums2 中的值在set中，则将此值加入到ret中
			delete(set, v)       //如果不加这个，那么 nums2中 如果有重复的值的话会重复 for range ，与nums1进行比较，重复加入到set中
			// 所以索引到一个后就给它从 nums1 中删掉
		}
	}
	return ret
}

// value, ok := map[key]
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	ret := intersection(nums1, nums2)
	fmt.Println(ret)
}

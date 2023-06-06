package main

import (
	"fmt"
	"sort"
)

// 四数之和
func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue //对i剪枝
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { //从 i+1开始索引，这样的话就不会索引到 i的前一次的 j-1的值了
				continue //对 j 剪枝
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					ret = append(ret, []int{n1, n2, n3, n4})
					//剪枝
					for l < r && n3 == nums[l] { //总结了下面的 这部分，很好！！！
						l++
					}
					for l < r && n4 == nums[r] {
						r--
					}
				} else if n1+n2+n3+n4 < target {
					l++
				} else {
					r--
				}
			}

		}
	}
	return ret
}

// 代码随想录
func fourSum2(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		// if n1 > target { // 不能这样写,因为可能是负数
		// 	break
		// }
		if i > 0 && n1 == nums[i-1] { // 对nums[i]去重 i-1，因为 i+1 的话，就缺少了 j 的索引 ，j 也得遍历那个数啊
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 对nums[j]去重
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l+1] { // 去重 注意此处是 l+1 ，跟前面的不一样  ，一个从左往右找，一个从右往左找，不存在前面那个问题
						l++
					}
					for l < r && n4 == nums[r-1] { // 去重
						r--
					}

					//错误的 去重操作没有做好
					//[-2,-1,-1,1,1,2,2]
					//0
					//输出：[[-2,-1,1,2],[-2,-1,1,2],[-1,-1,1,1]]
					//正确：[[-2,-1,1,2],[-1,-1,1,1]]
					//for l < r && l > j+1 && n3 == nums[l-1] { // 去重
					//	l++
					//}
					//for l < r && r < len(nums)-1 && n4 == nums[r+1] { // 去重
					//	r--
					//}

					// 找到答案时,双指针同时靠近
					r--
					l++
				}
			}
		}
	}
	return res
}
func main() {

	ret := fourSum2([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(ret)
}

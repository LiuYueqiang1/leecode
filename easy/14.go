package main

import "fmt"

// 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				fmt.Println(strs[0][:i])
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func main() {
	strs := make([]string, 0, 10)
	strs = []string{"flower", "flow", "flowight", "flowt"}
	longestCommonPrefix(strs)
}

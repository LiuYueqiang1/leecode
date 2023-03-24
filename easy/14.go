package main

// 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	var result string
	lowl := Lowlen(strs)
loop:
	for i := 0; i < lowl; i++ {
		str := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != str {
				break loop
			}
		}
		result += string(str)
	}
	return result
}

func Lowlen(s []string) int {
	lowlen := len(s[0])
	for i := 1; i < len(s); i++ {
		if len(s[i]) < lowlen {
			lowlen = len(s[i])
		}
	}
	return lowlen
}
func main() {
	strs := make([]string, 0, 10)
	strs = []string{"flower", "flow", "flight"}
	longestCommonPrefix(strs)
}

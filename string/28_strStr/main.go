package main

import "fmt"

// 方法一:前缀表使用减1实现

// getNext 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNext(next []int, s string) {
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] { // 如果当前字符不等于前一位相等前后缀的下一个字符，回退到上一个相等前后缀的位置
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] { // 如果当前字符等于前一位相等前后缀的下一个字符，j加1，指向当前相等前后缀的末尾
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}

/*
这部分代码是getNext函数，它的作用是构造前缀表next。next[i]表示字符串前缀s[0:i]中最长的相等前后缀长度，
也就是说，如果next[i]=k，那么s[0:k]就是s[0:i]的最长相等前后缀。
具体实现方法是，从第二个字符开始，依次遍历字符串，对于每个字符，我们通过回退j的方法找到它的最长相等前后缀。
j的初始值为-1，next[0]=-1，表示第一个字符没有相等前后缀。如果当前字符与前一位相等前后缀的下一个字符相同，
说明相等前后缀可以扩展，此时j加1，指向当前相等前后缀的末尾；
如果不同，说明相等前后缀不能再扩展，需要回退到上一个相等前后缀的位置，也就是next[j]处，继续判断是否相同。
如果j回退到-1还没找到相等前后缀，则表示当前字符前面没有相等前后缀，next[i]=-1。
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle)) // 构造前缀表
	getNext(next, needle)
	j := -1 // 模式串的起始位置next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 回退到下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1 // 返回匹配串的起始位置
		}
	}
	return -1
}

/*
这部分代码是主函数strStr，它的作用是在文本串haystack中查找模式串needle。
如果模式串为空，则返回0。否则，根据模式串构造前缀表next，并初始化模式串的起始位置j为-1。
从文本串的第一个字符开始遍历，对于每一个字符，如果它与模式串当前匹配的字符不同，
那么就回退j到下一个匹配点，即next[j]处，继续匹配；如果匹配成功，j加1，指向模式串当前匹配的字符的下一个位置。
如果j指向了模式串的末尾，说明已经找到了匹配串，返回匹配串的起始位置；
否则，继续遍历文本串。如果遍历完文本串还没有找到匹配串，则返回-1。
*/
func main() {
	s1 := "aabaabaafa"
	s2 := "aabaaf"
	n := strStr(s1, s2)
	fmt.Println(n)
}

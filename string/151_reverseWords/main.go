package main

import "fmt"

// 给定一个字符串，逐个翻转字符串中的每个单词。
//
// 示例 1：
// 输入: "the sky is blue"
// 输出: "blue is sky the"
//
// 示例 2：
// 输入: "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 示例 3：
// 输入: "a good   example"
// 输出: "example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
func reverseWords(s string) string {
	//slow 是过滤后的，fast是原来的指针
	slowindex, fastindex := 0, 0
	b := []byte(s)
	//删除头部空格
	for len(b) > 0 && b[fastindex] == ' ' && fastindex < len(b) {
		fastindex++ //将 fastindex 移到第一个字符位
	}
	//删除单词直接的冗余空格
	for ; fastindex < len(b); fastindex++ {
		if fastindex-1 > 0 && b[fastindex-1] == b[fastindex] && b[fastindex] == ' ' {
			continue
		}
		b[slowindex] = b[fastindex]
		//最后一次slow=12
		slowindex++
		//slow=13
		//此时b[12]=' ' , b[13]=' '
	}
	// 删除末尾的空格
	// 因为最后的时候slow++
	// []切片是左包含右不包含
	if slowindex-1 > 0 && b[slowindex-1] == ' ' {
		b = b[:slowindex-1]
	} else {
		b = b[:slowindex]
	}
	//反转整个字符串
	reverse(&b, 0, len(b)-1)
	//反转各个单词中的字符
	//以‘ ’为界
	//双指针

	//1、将i，j=0
	i := 0
	for i < len(b) {
		//2、将j=单词前的‘ ’
		j := i //定一个第一层for循环内全局的，后面用
		for ; j < len(b) && b[j] != ' '; j++ {
			// 将j移动到‘ ’处
		}
		//3、i，j反转
		reverse(&b, i, j-1)
		//4、i=j 也来到‘ ’
		i = j
		//5、i+1 来到单词中的第一个字母的位置
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
func main() {
	s := "hello world!   "
	res := reverseWords(s)
	fmt.Println(res)
}

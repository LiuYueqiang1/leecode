package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 三种情况
// 1、左边多余了 { { { ( ) } }
// 2、中间不匹配
// 3、右边多余了
func isValid(s string) bool {
	hash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] { // 条件3 条件2
			// 只要是匹配，栈的长度一定>0 且在这个else if 执行
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0 //条件 1  如果最后栈没有清空的话
}
func isValid2(s string) bool {
	hash := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[v] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
func main() {
	a := "{{[()]}"
	//123 91 40 41 93 125
	fmt.Println(a[1]) //91
	fmt.Println(')')
	s := []int{1, 2, 3, 4}
	fmt.Println(s[1]) //切片中的第几个数
	q := isValid(a)
	fmt.Println(q)
}

package main

import (
	"container/list"
	"fmt"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
//示例 1：
//输入：s = "()"
//输出：true
//输入：s = "()[]{}"
//输出：true
//输入：s = "(]"
//输出：false

var tokenMapping = map[byte]byte{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := list.New()
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			stack.PushBack(s[i])
		case ')', '}', ']':
			if stack.Len() > 0 {
				character := stack.Remove(stack.Back()).(byte)
				if tokenMapping[s[i]] != character {
					return false
				}
			} else {
				return false
			}

		}
	}
	return stack.Len() == 0
}
func main() {
	a := isValid("[][]{}()")
	fmt.Println(a)
}

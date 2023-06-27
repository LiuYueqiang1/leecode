package main

import (
	"fmt"
)

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
type stack struct {
	TopStack    int
	BottomStack int
	ValidArr    []int32
}

func isValid(s string) bool {
	st := stack{
		TopStack:    -1,
		BottomStack: 0,
		ValidArr:    make([]int32, 0),
	}
	for _, v := range s {
		if v == 123 || v == 91 || v == 40 {
			st.ValidArr = append(st.ValidArr, v)
			st.TopStack++
		} else if v == 41 || v == 93 || v == 125 {

		}
	}

	return true
}

func main() {
	a := "{[()]}"
	//123 91 40 41 93 125
	for _, v := range a {
		fmt.Printf(" %d", v)
	}

}

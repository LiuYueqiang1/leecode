package main

import "fmt"

// 383. 赎金信
// 给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
// 判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。
// 如果可以构成，返回 true ；否则返回 false。
func canConstruct2(ransomNote string, magazine string) bool {
	//ab ba 是可以的
	m := make(map[int32]bool, 0)
	count := 0
	for _, v := range ransomNote {
		count++
		m[v] = true //因为是aa ，所map的缺点就是不能全部索引，只有map[a]
	}
	for _, v := range magazine {
		m[v] = false
	}
	for _, v := range m {
		if v == false {
			count--
		}
	}
	return count == 0
}
func canConstruct(ransomNote string, magazine string) bool {
	set := make([]int, 36)
	for _, v := range ransomNote {
		set[v-'a']++
	}
	for _, v := range magazine {
		set[v-'a']--
		if set[v-'a'] < 0 {
			return false
		}
	}
	return true
}
func main() {
	//a := "string"
	//for _, v := range a {
	//	fmt.Printf("%T\n", v) //int32
	//}
	ret := canConstruct("aa", "ab")
	fmt.Println(ret)
}

package main

import "fmt"

//map[KeyType]ValueType   在go语言中用它来表示set  哈希有大用

// value, ok := map[key]
// map有 delete 函数  delete(map,key)

// make(map[KeyType]ValueType, [cap])
// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap) //map[小明:100 张三:90]
	// 如果key存在 ok 为true,v为对应的值；不存在 ok 为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

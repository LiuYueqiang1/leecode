package main

import "fmt"

func main() {
	m := make(map[int]bool)
	m[4] = true
	fmt.Println(m[4])
	fmt.Println(!m[4])
	fmt.Println(!m[6])
	fmt.Println()
	if m[5] {
		fmt.Println(1)
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"+", "-", "*", "/"}
	for _, v := range s {
		fmt.Println(v)
	}
	t := "+"

	fmt.Println(strconv.Atoi(t))
}

package main

import "math/big"

// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
func addBinary(a string, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)

	ai.Add(ai, bi)
	return ai.Text(2)
}
func main() {
	addBinary("11", "1")

}

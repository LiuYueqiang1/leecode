package main

import "fmt"

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 反转树
func invertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = node.Right, node.Left
	invertTree(node.Left) // 如果空了，此函数就返回
	invertTree(node.Right)
	return node
}

// 前序遍历
func orderTree(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: 4,
	}
	left1 := &TreeNode{
		Val: 2,
	}
	right1 := &TreeNode{
		Val: 7,
	}
	left2 := &TreeNode{
		Val: 1,
	}
	right2 := &TreeNode{
		Val: 3,
	}
	left3 := &TreeNode{
		Val: 6,
	}
	right3 := &TreeNode{
		Val: 9,
	}
	root.Left = left1
	root.Right = right1
	left1.Left = left2
	left1.Right = right2
	right1.Left = left3
	right1.Right = right3

	res := orderTree(root)
	fmt.Printf(" %d", res)
	fmt.Println()
	invertTree(root)
	res2 := orderTree(root)
	fmt.Printf(" %d", res2)
}

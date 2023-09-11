package main

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个函数、调用这个函数即可
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var pretree func(node *TreeNode)
	pretree = func(node *TreeNode) {
		if node != nil {
			ret = append(ret, node.Val)
			pretree(node.Left)
			pretree(node.Right)
		}
	}
	pretree(root)
	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	var res = make([]int, 0)
	var traversal func(node *TreeNode)
	// 定义一个函数
	traversal = func(node *TreeNode) {
		if node != nil {
			res = append(res, node.Val)
			traversal(node.Left)
			traversal(node.Right)
		}
	}

	traversal(root)
	return res
}

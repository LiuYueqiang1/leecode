package main

import "fmt"

// 对称二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func defs(left *TreeNode, right *TreeNode) bool {
	// 1、如果左右节点都为空，则返回true
	if left == nil && right == nil {
		return true
	}
	// 2、如果左右节点有一个为空
	if left == nil || right == nil {
		return false
	}
	// 3、如果左节点不等于右节点
	if left.Val != right.Val {
		return false
	}
	// 左边的左节点、右边的右节点；左边的右节点、右边的左节点
	return defs(left.Left, right.Right) && defs(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
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
	// 创建一个示例二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	// 调用isSymmetric函数判断二叉树是否对称
	result := isSymmetric(root)
	fmt.Println(result) // 输出true，示例二叉树是对称的
}

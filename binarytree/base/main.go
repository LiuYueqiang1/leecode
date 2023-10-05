package main

/*
递归三要素：
确定递归函数的参数和返回值： 确定哪些参数是递归的过程中需要处理的，那么就在递归函数里加上这个参数，
并且还要明确每次递归的返回值是什么进而确定递归函数的返回类型。

确定终止条件： 写完了递归算法, 运行的时候，经常会遇到栈溢出的错误，就是没写终止条件或者终止条件写的不对，
操作系统也是用一个栈的结构来保存每一层递归的信息，如果递归没有终止，操作系统的内存栈必然就会溢出。

确定单层递归的逻辑： 确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程。
*/

//// 前序遍历[先输root结点，然后再输出左子树，然后再输出右子树]
//func PreOrder(hero *Hero) {
//	if hero != nil {
//		fmt.Printf("no=%d name=%s\n", hero.No, hero.Name)
//		PreOrder(hero.Left)
//		PreOrder(hero.Right)
//	}
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 根左右
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
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

// 中序遍历 左根右
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 后续遍历 左右根
func postorderTraversal(root *TreeNode) (res []int) {
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
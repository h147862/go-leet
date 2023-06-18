package BinaryTreeInorderTraversal

// https://leetcode.com/problems/binary-tree-inorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	InOrder(root, &res)
	return res
}

func InOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if &node.Left != nil {
		InOrder(node.Left, res)
	}
	*res = append(*res, node.Val)
	if &node.Right != nil {
		InOrder(node.Right, res)
	}
}

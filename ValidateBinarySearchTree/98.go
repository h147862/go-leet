package ValidateBinarySearchTree

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).

// A valid BST is defined as follows:

// The left
// subtree
//  of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

// Input: root = [2,1,3]
// Output: true

// Input: root = [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeQueue struct {
	item []*TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := []int{}
	InOrder(root, &res)
	max := math.Inf(-1)
	for _, v := range res {
		if float64(v) <= max {
			return false
		}
		max = float64(v)
	}
	return true
}

func InOrder(node *TreeNode, res *[]int) {
	if node.Left != nil {
		InOrder(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		InOrder(node.Right, res)
	}
}

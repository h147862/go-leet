package FindModeinBinarySearchTree

// https://leetcode.com/problems/find-mode-in-binary-search-tree/
// Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.

// If the tree has more than one mode, return them in any order.

// Assume a BST is defined as follows:

// The left subtree of a node contains only nodes with keys less than or equal to the node's key.
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
// Both the left and right subtrees must also be binary search trees.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	count := make(map[int]int)
	PreOrder(root, &count)
	max := 0
	res := []int{}
	for k, v := range count {
		if v > max {
			res = nil
			max = v
			res = append(res, k)
		} else if v == max {
			max = v
			res = append(res, k)
		}
	}
	return res
}

func PreOrder(root *TreeNode, count *map[int]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		PreOrder(root.Left, count)
	}
	_, ok := (*count)[root.Val]
	if ok {
		(*count)[root.Val] += 1
	} else {
		(*count)[root.Val] = 1
	}
	if root.Right != nil {
		PreOrder(root.Right, count)
	}
}

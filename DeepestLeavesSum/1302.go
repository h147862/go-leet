package DeepestLeavesSum

// https://leetcode.com/problems/deepest-leaves-sum/

// Given the root of a binary tree, return the sum of values of its deepest leaves.

// Example 1:
// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// Output: 15
// Example 2:

// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	return DeepestLeavesSum(root)
}

func DeepestLeavesSum(root *TreeNode) int {
	count := make(map[int]int)
	height := -1
	DFSInorder(root, &count, height)
	max := 0
	for h, _ := range count {
		if h >= max {
			max = h
		}
	}
	return count[max]
}

func DFSInorder(node *TreeNode, count *map[int]int, heght int) {
	if node == nil {
		return
	}
	heght += 1
	DFSInorder(node.Left, count, heght)
	_, ok := (*count)[heght]
	if !ok {
		(*count)[heght] = node.Val
	} else {
		(*count)[heght] += node.Val
	}
	DFSInorder(node.Right, count, heght)
}

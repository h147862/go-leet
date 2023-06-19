package main

import (
	"fmt"
	v "go-leet/DeepestLeavesSum"
)

func main() {

	root := &v.TreeNode{
		Val: 1,
		Left: &v.TreeNode{
			Val: 2,
			Left: &v.TreeNode{
				Val: 4,
				Left: &v.TreeNode{
					Val: 7,
				},
			},
			Right: &v.TreeNode{
				Val: 5,
			},
		},
		Right: &v.TreeNode{
			Val: 3,
			Left: &v.TreeNode{
				Val: 6,
				Right: &v.TreeNode{
					Val: 8,
				},
			},
		},
	}
	res := v.DeepestLeavesSum(root)
	fmt.Print(res)
}

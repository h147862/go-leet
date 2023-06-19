package main

import (
	"fmt"
	v "go-leet/FindModeinBinarySearchTree"
)

func main() {

	root := v.TreeNode{
		Val:  1,
		Left: nil,
		Right: &v.TreeNode{
			Val: 2,
			Left: &v.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	res := make(map[int]int)
	v.PreOrder(&root, &res)
	fmt.Print(res)
}

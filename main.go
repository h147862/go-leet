package main

import (
	"fmt"
	v "go-leet/ValidateBinarySearchTree"
)

func main() {

	root := v.TreeNode{
		Val: 0,
		Left: &v.TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
	}
	res := []int{}
	v.InOrder(&root, &res)
	fmt.Print(res)
}

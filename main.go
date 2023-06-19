package main

import (
	"fmt"
	v "go-leet/MaximumBinaryTree"
)

func main() {

	nums := []int{3, 2, 1, 6, 0, 5}
	node := v.ConstructMaximumBinaryTree(nums)
	res := v.BFS(node)
	fmt.Print(res)
}

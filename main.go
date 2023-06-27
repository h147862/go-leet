package main

import (
	"fmt"
	m "go-leet/Subsets"
)

func main() {

	nums := []int{1, 2, 3}
	res := [][]int{[]int{}}
	cur := []int{}
	m.Subsets(nums, &res, &cur)
	fmt.Println(res)
}

package main

import (
	"fmt"
	m "go-leet/CombinationSum"
	"sort"
)

func main() {
	nums := []int{2, 9, 6, 7}
	target := 7
	cur := []int{}
	res := [][]int{}
	idx := 0
	sort.IntSlice(nums).Sort()
	m.CombinationSum(nums, target, &cur, &res, idx)
	fmt.Println(res)
}

package main

import (
	"fmt"
	m "go-leet/SubsetsII"
	"sort"
)

func main() {
	nums := []int{4, 4, 4, 1, 4}
	cur := []int{}
	res := [][]int{}
	used := []bool{}
	exist := make(map[string]bool)
	sort.Ints(nums)
	for idx := 0; idx < len(nums); idx++ {
		used = append(used, false)
	}
	m.SubsetsWithDup(nums, &used, &exist, &cur, &res, 0)
	fmt.Println(res)
}

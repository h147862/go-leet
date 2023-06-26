package main

import (
	"fmt"
	m "go-leet/Permutations"
)

func main() {
	nums := []int{1, 2, 3}
	used := make(map[int]bool)
	cnt := 0
	result := [][]int{}
	tmp := []int{}
	m.Permute(nums, used, cnt, &result, &tmp)
	fmt.Println(result)
}

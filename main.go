package main

import (
	"fmt"
	m "go-leet/Combinations"
)

func main() {
	nums := []int{1, 2, 3, 5, 7}
	k := 3
	used := make(map[int]map[int]bool)
	for i := 0; i < k; i++ {
		used[i] = make(map[int]bool)
	}
	cnt := 0
	result := [][]int{}
	tmp := []int{}
	m.Combine(nums, k, used, &tmp, cnt, &result)
	fmt.Println(result)
}

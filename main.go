package main

import (
	"fmt"
	m "go-leet/CombinationSumII"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	sort.Ints(candidates)
	used := make([]int, len(candidates))
	cur := []int{}
	res := [][]int{}
	exist := make(map[string]bool)
	m.Combi(candidates, target, &used, &cur, &res, &exist, 0)
	fmt.Println(res)
}

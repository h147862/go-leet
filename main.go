package main

import (
	"fmt"
	m "go-leet/CombinationSumII"
	"sort"
)

func main() {
	candidates := []int{1, 1, 1, 3, 3, 5}
	target := 8
	sort.Ints(candidates)
	cur := []int{}
	res := [][]int{}
	m.Combi(candidates, target, &cur, &res, 0)
	fmt.Println(res)
}

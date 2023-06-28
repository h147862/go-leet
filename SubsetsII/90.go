// https://leetcode.com/problems/subsets-ii/
// Given an integer array nums that may contain duplicates, return all possible
// subsets
//  (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:

// Input: nums = [1,2,2]
// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

package SubsetsII

import (
	"sort"
	"strconv"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	cur := []int{}
	res := [][]int{[]int{}}
	used := []bool{}
	exist := make(map[string]bool)
	sort.Ints(nums)
	for idx := 0; idx < len(nums); idx++ {
		used = append(used, false)
	}
	if len(nums) == 0 {
		return res
	}
	SubsetsWithDup(nums, &used, &exist, &cur, &res, 0)
	return res
}

func SubsetsWithDup(nums []int, used *[]bool, exist *map[string]bool, cur *[]int, res *[][]int, start int) {
	for idx := start; idx < len(nums); idx++ {
		if (*used)[idx] {
			continue
		}
		*cur = append(*cur, nums[idx])
		k := getKey(*cur)
		if b, ok := (*exist)[k]; !b || !ok {
			(*exist)[k] = true
		} else {
			*cur = (*cur)[:len(*cur)-1]
			continue
		}
		(*used)[idx] = true
		*res = append(*res, append([]int{}, *cur...))
		SubsetsWithDup(nums, used, exist, cur, res, idx+1)
		*cur = (*cur)[:len(*cur)-1]
		(*used)[idx] = false
	}
}

func getKey(cur []int) string {
	str := []string{}
	for _, num := range cur {
		str = append(str, strconv.Itoa(num))
	}
	k := strings.Join(str, "-")
	return k
}

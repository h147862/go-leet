// https://leetcode.com/problems/combination-sum-ii/
// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.

// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]

// Example 2:

// Input: candidates = [2,5,2,1,2], target = 5
// Output:
// [
// [1,2,2],
// [5]
// ]

package CombinationSumII

import (
	"sort"
	"strconv"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	used := make([]int, len(candidates))
	cur := []int{}
	res := [][]int{}
	exist := make(map[string]bool)
	Combi(candidates, target, &used, &cur, &res, &exist, 0)
	return res
}

func Combi(nums []int, target int, used *[]int, cur *[]int, res *[][]int,
	exist *map[string]bool, start int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *cur...))
	}
	for idx := start; idx < len(nums); idx++ {
		if (*used)[idx] != 0 {
			continue
		}
		if target-nums[idx] >= 0 {
			*cur = append(*cur, nums[idx])
			k := getKey(*cur)
			if b, ok := (*exist)[k]; !b || !ok {
				(*exist)[k] = true
				(*used)[idx] += 1
				Combi(nums, target-nums[idx], used, cur, res, exist, idx+1)
				(*used)[idx] -= 1
				*cur = (*cur)[:len(*cur)-1]
			} else {
				*cur = (*cur)[:len(*cur)-1]
				continue
			}

		} else {
			return
		}
	}
}

func getKey(cur []int) string {
	str := []string{}
	for _, v := range cur {
		str = append(str, strconv.Itoa(v))
	}
	key := strings.Join(str, "-")
	return key
}

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
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	cur := []int{}
	res := [][]int{}
	Combi(candidates, target, &cur, &res, 0)
	return res
}

func Combi(nums []int, target int, cur *[]int, res *[][]int, start int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *cur...))
	}
	for idx := start; idx < len(nums); idx++ {
		if idx > start && nums[idx] == nums[idx-1] {
			continue
		}
		if target-nums[idx] >= 0 {
			*cur = append(*cur, nums[idx])
			Combi(nums, target-nums[idx], cur, res, idx+1)
			*cur = (*cur)[:len(*cur)-1]
		} else {
			return
		}
	}
}

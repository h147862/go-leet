// https://leetcode.com/problems/permutations/

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// Example 2:

// Input: nums = [0,1]
// Output: [[0,1],[1,0]]

// Example 3:

// Input: nums = [1]
// Output: [[1]]
package Permutations

// func permute(nums []int) [][]int {

// }

func permute(nums []int) [][]int {
	used := make(map[int]bool)
	cnt := 0
	result := [][]int{}
	tmp := []int{}
	Permute(nums, used, cnt, &result, &tmp)
	return result
}

func Permute(nums []int, used map[int]bool, cnt int, res *[][]int, tmp *[]int) {
	if cnt == len(nums) {
		newSlice := make([]int, len(*tmp))
		copy(newSlice, *tmp)
		*res = append(*res, newSlice)
		tmp = nil
	} else {
		for _, v := range nums {
			b, ok := used[v]
			if !ok || !b {
				*tmp = append(*tmp, v)
				used[v] = true
			} else {
				continue
			}
			Permute(nums, used, cnt+1, res, tmp)
			used[v] = false
			*tmp = (*tmp)[:cnt]
		}
	}
}

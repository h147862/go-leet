// https://leetcode.com/problems/combinations/

// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

// You may return the answer in any order.

// Example 1:

// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
// Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

// Example 2:

// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is 1 choose 1 = 1 total combination.

package Combinations

func combine(n int, k int) [][]int {
	nums := []int{}
	for idx := 1; idx <= n; idx++ {
		nums = append(nums, idx)
	}
	used := make(map[int]bool)
	cnt := 0
	result := [][]int{}
	tmp := []int{}
	Combine(nums, k, used, &tmp, cnt, &result)
	return result
}

func Combine(nums []int, k int, used map[int]bool, tmp *[]int, cnt int, res *[][]int) {
	if cnt == k {
		*res = append(*res, append([]int{}, *tmp...))
		return
	}
	for idx, n := range nums {
		isUsed, hasValue := used[n]
		if !isUsed || !hasValue {
			*tmp = append(*tmp, n)
			used[n] = true
		} else {
			continue
		}
		Combine(nums[idx+1:], k, used, tmp, cnt+1, res)
		*tmp = (*tmp)[:cnt]
		used[n] = false
	}
}

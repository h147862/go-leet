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
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return [][]int{}
}

func Combine(nums []int, k int, used map[int]map[int]bool, tmp *[]int, cnt int, res *[][]int) {
	if cnt == k {
		*res = append(*res, append([]int{}, *tmp...))
		return
	}
	for _, n := range nums {
		isUsed, hasValue := used[cnt][n]
		if !isUsed || !hasValue {
			if cnt == 0 {
				*tmp = append(*tmp, n)
				used[cnt][n] = true
			} else {
				b, ok := used[cnt-1][n]
				if !b || !ok {
					*tmp = append(*tmp, n)
					used[cnt][n] = true
				} else {
					continue
				}
			}
			Combine(nums, k, used, tmp, cnt+1, res)
			*tmp = (*tmp)[:cnt]
		}
	}
}

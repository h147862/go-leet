// https://leetcode.com/problems/top-k-frequent-elements/
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	return Fre(nums, k)
}

func Fre(nums []int, k int) []int {
	result := []int{}
	cnt := make(map[int]int)
	for _, n := range nums {
		if _, exist := cnt[n]; !exist {
			cnt[n] = 1
		} else {
			cnt[n]++
		}
	}
	for idx := 0; idx < k; idx++ {
		max := 0
		target := -1

		for k, v := range cnt {
			if v >= max {
				max = v
				target = k
			}
		}
		result = append(result, target)
		delete(cnt, target)
	}
	return result
}

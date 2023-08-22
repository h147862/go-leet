// https://leetcode.com/problems/product-of-array-except-self/
// 238. Product of Array Except Self

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

package ProductOfArrayExceptSelf

func productExceptSelf(nums []int) []int {

	return P(nums)
}

func P(nums []int) []int {
	left := []int{}
	right := []int{}
	res := []int{}
	for idx, _ := range nums {
		if idx == 0 {
			left = append(left, 1)
			right = append(right, 1)
		} else {
			left = append(left, left[idx-1]*nums[idx-1])
			right = append(right, right[idx-1]*nums[len(nums)-idx])
		}
	}
	for idx, _ := range nums {
		v := left[idx] * right[len(nums)-idx-1]
		res = append(res, v)
	}
	return res
}

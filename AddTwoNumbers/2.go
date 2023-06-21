package AddTwoNumbers

// https://leetcode.com/problems/add-two-numbers/

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]

// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return AddTwoNumbers(l1, l2)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := []ListNode{}
	carry := 0
	curVal := 0
	for true {
		if l1 != nil && l2 != nil {
			v := l1.Val + l2.Val + carry
			carry = 0
			if v >= 10 {
				carry = v / 10
				curVal = v % 10
			} else {
				curVal = v
			}
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			v := l1.Val + carry
			carry = 0
			if v >= 10 {
				carry = v / 10
				curVal = v % 10
			} else {
				curVal = v
			}
			l1 = l1.Next
		} else if l2 != nil {
			v := l2.Val + carry
			carry = 0
			if v >= 10 {
				carry = v / 10
				curVal = v % 10
			} else {
				curVal = v
			}
			l2 = l2.Next
		} else {
			if carry != 0 {
				res = append(res, ListNode{Val: carry})
			}
			break
		}
		res = append(res, ListNode{Val: curVal})
	}
	if len(res) == 1 {
		return &res[0]
	} else {
		for idx := 0; idx < len(res)-1; idx++ {
			res[idx].Next = &res[idx+1]
		}
		return &res[0]
	}
}

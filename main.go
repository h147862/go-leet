package main

import (
	"fmt"
	m "go-leet/AddTwoNumbers"
)

func main() {

	// l1 := m.ListNode{
	// 	Val: 2,
	// 	Next: &m.ListNode{
	// 		Val: 4,
	// 		Next: &m.ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	// l2 := m.ListNode{
	// 	Val: 5,
	// 	Next: &m.ListNode{
	// 		Val: 6,
	// 		Next: &m.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	arr1 := []int{2, 4, 9}
	arr2 := []int{5, 6, 4, 9}

	res1 := []m.ListNode{}
	res2 := []m.ListNode{}

	for _, v := range arr1 {
		res1 = append(res1, m.ListNode{Val: v})
	}
	for _, v := range arr2 {
		res2 = append(res2, m.ListNode{Val: v})
	}

	for idx := 0; idx < len(res1)-1; idx++ {
		res1[idx].Next = &res1[idx+1]
	}
	for idx := 0; idx < len(res2)-1; idx++ {
		res2[idx].Next = &res2[idx+1]
	}

	node := m.AddTwoNumbers(&res1[0], &res2[0])
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

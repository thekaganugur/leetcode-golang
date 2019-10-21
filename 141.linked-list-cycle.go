/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

// Method 1: Floyd’s Cycle-Finding O(n), 4ms, 3.8mb
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return false
}

// Method 2: Hash Table O(n), 8ms, 5mb
// func hasCycle(head *ListNode) bool {
// 	m := make(map[*ListNode]struct{})

// 	for head != nil {
// 		if _, ok := m[head]; ok {
// 			return true
// 		}
// 		m[head] = struct{}{}
// 		head = head.Next
// 	}

// 	return false
// }

// @lc code=end

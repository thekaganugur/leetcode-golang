/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
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

// Method 1: Two pointers O(n+m), ms, mb
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	var linkedA, linkedB bool

	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa, pb = pa.Next, pb.Next

		if pa == nil && !linkedA {
			pa = headB
			linkedA = true
		}
		if pb == nil && !linkedB {
			pb = headA
			linkedB = true
		}
	}

	return nil
}

// Method 2: Hash table O(n+m), 56ms, 8.6mb
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	m := make(map[*ListNode]struct{})

// 	for headA != nil {
// 		m[headA] = struct{}{}
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if _, ok := m[headB]; ok {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}

// 	return nil
// }

// @lc code=end

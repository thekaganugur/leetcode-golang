/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
package main

// Method 1: Recursive
func searchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1
	return searchInsertUtil(nums, target, l, h)
}

func searchInsertUtil(nums []int, target, l, h int) int {
	if l > h {
		return l
	}
	m := (l + h) / 2

	if target == nums[m] {
		return m
	} else if target < nums[m] {
		return searchInsertUtil(nums, target, l, m-1)
	} else {
		return searchInsertUtil(nums, target, m+1, h)
	}
}

// Method 2: Iterative
// func searchInsert(nums []int, target int) int {
// 	l, h := 0, len(nums)-1

// 	for l <= h {
// 		m := (l + h) / 2

// 		if target > nums[m] {
// 			l = m + 1
// 		} else if target < nums[m] {
// 			h = m - 1
// 		} else {
// 			return m
// 		}
// 	}

// 	return l
// }

// @lc code=end

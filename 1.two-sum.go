/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		ii, ok := m[v]
		if ok {
			return []int{ii, i}
		}
		m[target-v] = i
	}
	return nil
}

// @lc code=end

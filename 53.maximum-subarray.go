/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
package main

func maxSubArray(nums []int) int {
	lm, gm := nums[0], nums[0]

	for _, v := range nums[1:] {
		if lm+v < v {
			lm = v
		} else {
			lm += v
		}

		if lm > gm {
			gm = lm
		}
	}

	return gm
}

// @lc code=end

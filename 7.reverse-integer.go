/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
package main

import "math"

func reverse(x int) (y int) {
	for x != 0 {
		popped := x % 10
		x /= 10

		tmp := y*10 + popped
		y = tmp

		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
	}
	return
}

// @lc code=end

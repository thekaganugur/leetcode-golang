/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
package main

import "math"

func isPalindrome(x int) bool {
	if math.Signbit(float64(x)) {
		return false
	}

	//take the reverse
	xx := x
	reversed := 0
	for xx != 0 {
		pop := xx % 10
		xx /= 10

		tmp := reversed*10 + pop
		reversed = tmp
	}

	if x == reversed {
		return true
	}

	return false
}

// @lc code=end

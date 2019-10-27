/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
package main

func convertToTitle(n int) string {
	b := make([]byte, 0)

	for n > 0 {
		n--
		b = append([]byte{byte('A' + n%26)}, b...)
		n /= 26
	}

	return string(b)
}

// @lc code=end

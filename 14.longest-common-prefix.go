/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
package main

import "strings"

// Approach 1: Vertical scanning (0 ms)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minlen := findMinLength(strs)
	var lcp strings.Builder

	for i := 0; i < minlen; i++ {
		firstStrCh := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if firstStrCh != strs[j][i] {
				return lcp.String()
			}
		}
		lcp.WriteByte(firstStrCh)
	}

	return lcp.String()
}

func findMinLength(strs []string) (m int) {
	m = len(strs[0])
	for _, str := range strs {
		if len(str) < m {
			m = len(str)
		}
	}
	return
}

// Approach 1: Horizontal scanning (4 ms)
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	lcp := strs[0]

// 	for _, str := range strs {
// 		tmp := ""

// 		for j := 0; j < len(lcp) && j < len(str); j++ {
// 			lcpCh, strCh := lcp[j], str[j]

// 			if lcpCh == strCh {
// 				tmp += string(lcpCh)
// 			} else {
// 				break
// 			}
// 		}

// 		lcp = tmp
// 	}

// 	return lcp
// }

// @lc code=end

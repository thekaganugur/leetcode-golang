/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
package main

import "math"

// Method 1: One Pass (4ms)
func maxProfit(prices []int) (profit int) {
	min := math.MaxInt64
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > profit {
			profit = v - min
		}
	}
	return
}

// Method 2: Brute force, (192ms)
// func maxProfit(prices []int) (max int) {
// 	for i := len(prices) - 1; i > -1; i-- {
// 		for j := i - 1; j > -1; j-- {
// 			dif := prices[i] - prices[j]
// 			if dif > max {
// 				max = dif
// 			}
// 		}
// 	}
// 	return
// }

// @lc code=end

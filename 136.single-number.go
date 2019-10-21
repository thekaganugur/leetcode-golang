/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
package main

//  Bit Manipulation XOR, 4ms, 4.7mb
func singleNumber(nums []int) (res int) {
	for _, v := range nums {
		res ^= v
	}
	return
}

// 12ms, 5.8mb
// func singleNumber(nums []int) (sum int) {
// 	m := make(map[int]int)

// 	for _, v := range nums {
// 		if m[v] == 0 {
// 			m[v]++
// 			sum += v
// 		} else {
// 			delete(m, v)
// 			sum -= v
// 		}
// 	}

// 	return
// }

// 12ms, 5.8mb
// func singleNumber(nums []int) (single int) {
// 	m := make(map[int]int)

// 	for _, v := range nums {
// 		m[v]++
// 		if m[v] == 2 {
// 			delete(m, v)
// 		}
// 	}
// 	for v := range m {
// 		single = v
// 	}

// 	return
// }

// @lc code=end

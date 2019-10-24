/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
package main

func twoSum(numbers []int, target int) (res []int) {
	m := make(map[int]int)

	for i, v := range numbers {
		if ii, ok := m[target-v]; ok {
			ii++
			res = append(res, ii)
			i++
			return append(res, i)
		}
		m[v] = i
	}

	return
}

// @lc code=end

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package main

// @lc code=start
// Method 1: Bottom up dymanic programming, O(n)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a := make([]int, n+1)
	a[1], a[2] = 1, 2

	for i := 3; i <= n; i++ {
		a[i] = a[i-2] + a[i-1]
	}

	return a[n]
}

// Method 2: Recursion with Memoization, O(n)
// func climbStairs(n int) int {
// 	memo := make([]int, n+1)
// 	return climbStairsUtil(0, n, memo)
// }
// func climbStairsUtil(i, n int, memo []int) int {
// 	if i > n {
// 		return 0
// 	}
// 	if i == n {
// 		return 1
// 	}
// 	if memo[i] > 0 {
// 		return memo[i]
// 	}
// 	memo[i] = climbStairsUtil(i+1, n, memo) + climbStairsUtil(i+2, n, memo)
// 	return memo[i]
// }

// Method 3: Brute Force, Recursion, O(2^n), Can't pass because of O(2^n)
// func climbStairs(n int) int {
// 	return climbStairsUtil(0, n)
// }
// func climbStairsUtil(i, n int) int {
// 	if i > n {
// 		return 0
// 	}
// 	if i == n {
// 		return 1
// 	}
// 	return climbStairsUtil(i+1, n) + climbStairsUtil(i+2, n)
// }

// @lc code=end

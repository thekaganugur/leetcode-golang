/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
package main

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum int
	var rr rune

	for i, r := range s {
		curr, prev := roman[r], roman[rr]
		rr = r
		sum += curr

		if i > 0 && prev < curr {
			sum -= prev * 2
		}
	}

	return sum
}

// Long approach
// func romanToInt(s string) int {
// 	sum, l := 0, len(s)

// 	for i := 0; i < l; i++ {
// 		r1 := roman[rune(s[i])]

// 		if i == l-1 {
// 			sum += r1
// 			return sum
// 		}

// 		r2 := roman[rune(s[i+1])]

// 		if r1 < r2 {
// 			sum -= r1
// 		} else {
// 			sum += r1
// 		}
// 	}

// 	return sum
// }

// @lc code=end

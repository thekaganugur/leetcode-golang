/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
package main

var m = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}
func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func isValid(str string) bool {
	s := make(stack, 0)

	for _, v := range str {
		l := len(s)

		switch v {
		case '(', '{', '[':
			s = s.Push(v)
		case ')', '}', ']':
			if l > 0 && m[v] == s[l-1] {
				s, _ = s.Pop()
				continue
			}
			return false
		}
	}

	return len(s) == 0
}

// @lc code=end

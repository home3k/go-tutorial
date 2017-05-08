package main

import "fmt"

/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

For "(()", the longest valid parentheses substring is "()", which has length = 2.

Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.
((())(()
*/
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	p, index := []rune{}, []int{}
	for i, r := range s {
		if r == '(' {
			p = append(p, '(')
			index = append(index, i)
		} else {
			pos := len(index) - 1
			if pos >= 0 && p[pos] == '(' {
				index = index[:pos]
				p = p[:pos]
			} else {
				index = append(index, i)
				p = append(p, ')')
			}
		}
	}
	if len(index) == 0 {
		return len(s)
	}
	max := 0
	i := 0
	for j, v := range index {
		var l int
		if j == 0 {
			l = v - i
		} else {
			l = v - i - 1
		}
		if l > max && l % 2 == 0 {
			max = l
		}
		i = v
	}
	l := len(s) - 1 - i
	if l > max {
		max = l
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses(")"))
	fmt.Println(longestValidParentheses("("))
	fmt.Println(longestValidParentheses("()"))
	fmt.Println(longestValidParentheses("())"))
	fmt.Println(longestValidParentheses("()()"))
	fmt.Println(longestValidParentheses("()())"))
	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses("((())(()"))
	fmt.Println(longestValidParentheses("()())()()"))
}

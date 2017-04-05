package main

import "fmt"

//isMatch("aa","a") → false
//isMatch("aa","aa") → true
//isMatch("aaa","aa") → false
//isMatch("aa", "a*") → true
//isMatch("aa", ".*") → true
//isMatch("ab", ".*") → true
//isMatch("aab", "c*a*b") → true
//
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	if lp == 0 {
		return ls == 0
	}
	// 一个字符的时候
	if lp == 1 {
		return ls == 1 && (p[0] == '.' || p[0] == s[0])
	}
	// > 1个字符的判断
	if p[1] != '*' {
		return ls != 0 &&  (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[1:])
	} else {
		// p[1] == *
		return (ls != 0 && (p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
}

func main()  {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aaa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("", "c*a*"))
}

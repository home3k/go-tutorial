package main

import "fmt"

/*
'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false

*/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return true
	}
	if len(s) == 0 {
		return p != "*"
	}
	ss, sp := 0, 0
	for ss < len(s) && sp < len(p) {
		if s[ss] == p[sp] || p[sp] == '?' {
			ss++
			sp++
		} else if p[sp] == '*' {
			return subMatch(s[ss:], fetchSub(p[sp+1:]))
		} else {
			return false
		}
	}
	if ss < len(s) {
		return false
	} else if sp >= len(p) {
		return true
	} else {
		for sp < len(p) {
			if p[sp] == '*' {
				sp++
			} else {
				return false
			}
		}
		return true
	}
}

func fetchSub(s string) []string {
	result := []string{}
	i := 0
	for x, v := range s {
		if v == '*' {
			t := s[i : x]
			if len(t) != 0 {
				result = append(result, t)
			}
			i=x+1
		}
	}
	t := s[i:]
	if len(t) != 0 {
		result = append(result, t)
	}
	return result
}

func subMatch(s string, sub []string) bool {
	if len(sub) == 0 {
		return true
	}
	if len(s) == 0 {
		return len(sub) == 0
	}
	ss, sSub, x := 0, 0, 0
	for ss < len(s) && sSub < len(sub) {
		for x < len(sub[sSub]) && ss<len(s) {
			if s[ss] == sub[sSub][x] || sub[sSub][x] == '?' {
				ss++
				x++
			} else {
				if x == 0 {
					ss++
				} else {
					x=0
				}
			}
		}
		if x >= len(sub[sSub]) {
			sSub++
			x = 0
		}

	}
	if ss >= len(s) && sSub >= len(sub) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", "?*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aab", "a*b*b"))
	fmt.Println(isMatch("aabb", "a*b*b"))
	fmt.Println(isMatch("missingtest", "mi*ing?s*t"))
	fmt.Println(isMatch("bdfasdfadfwefwfadfasd", "*fa??*we*?df*d"))
}

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	match := make(map[byte]bool)
	var start, max, count int
	l := len(s)
	for start < l {
		for index := start; index < l; index++ {
			val := s[index]
			if len(match) == 0 {
				start = index
			}
			if _, ok := match[val]; !ok {
				match[val] = true
				count++
			} else {
				if count > max {
					max = count
				}
				count = 0
				start++
				// quick fail
				if l-start <= max {
					return max
				}
				match = make(map[byte]bool)
				break
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dpwewkew"))
}

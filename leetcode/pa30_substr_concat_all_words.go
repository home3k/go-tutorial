package main

import "fmt"

/*

You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).

*/
func findSubstring(s string, words []string) []int {
	if s == "" || words == nil || len(words) == 0 {
		return []int{}
	}
	tl := len(words)
	l := len(words[0])
	if l == 0 {
		return []int{}
	}
	mapping := make(map[string]int)
	for _, word := range words {
		mapping[word] += 1
	}
	result := []int{}
	i := 0
	for ; i < len(s)-l*tl; i++ {
		if _, ok := mapping[s[i:i+l]]; !ok {
			continue
		}
		findIt := _findSubstr(s[i:i+l*tl], mapping, l, tl)
		if findIt {
			result = append(result, i)
		}
	}
	if i < len(s) {
		findIt := _findSubstr(s[i:], mapping, l, tl)
		if findIt {
			result = append(result, i)
		}
	}
	return result
}

func _findSubstr(s string, mapping map[string]int, step int, count int) bool {
	if len(s) < step*count {
		return false
	}
	head := s[0:step]
	if contain, ok := mapping[head]; ok && contain > 0 {
		if count == 1 {
			return true
		}
		mapping[head] -= 1
		result := _findSubstr(s[step:], mapping, step, count-1)
		mapping[head] += 1
		return result
	} else {
		return false
	}
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}

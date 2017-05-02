package main

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
	tl:=len(words)
	l:=len(words[0])
	if l == 0 {
		return []int{}
	}
	mapping:=make(map[string]bool)
	for _, word:=range words {
		mapping[word] = true
	}
	c:=0
	for i:=0;i<len(s)-l*tl;i+=l {
		if _, ok:=mapping[s[i:l]]; ok {
			
		}	
	}
}

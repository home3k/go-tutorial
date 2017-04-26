package main

import "fmt"

// 2 abc 3 def 4 ghi 5 jkl 6 mno 7 pgrs 8 tuv 9 wxyz
// Input:Digit string "23"
// abc def
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
/*
	234
	abc def ghi
	a b c   d e f   g h i
	a []
 */
var mapping = map[rune]string{
	'0': "",
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	letters:=[]string{}
	for _, r := range digits {
		if v, ok := mapping[r]; ok {
			letters = append(letters, v)
		}
	}
	if len(letters) == 0 {
		return []string{}
	}
	return _combi(letters)
}


func _combi(s[]string)[]string {
	result:=[]string{}
	if len(s) == 1 {
		for _, v:=range s[0] {
			result = append(result, string(v))
		}
		return result
	}
	next := _combi(s[1:])
	for _, v:= range s[0] {
		for _, c:= range next{
			result = append(result, string(v)+c)
		}
	}
	return result
}


func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
}

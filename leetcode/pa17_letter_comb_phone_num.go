package main

// 2 abc 3 def 4 ghi 5 jkl 6 mno 7 pgrs 8 tuv 9 wxyz
// Input:Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
var mapping = map[rune]string{
	'0': "",
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pgrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	for _, r := range digits {

	}
}

func combi(s1 string, s2 string, result []string) {
	if len(s1) == 1 && len(s2) == 1 {
		result = append(result, s1+s2)
	} else if len(s1) == 1 {
		result = append(result, s1+string(s2[0]))
		combi(s1, s2[1:], result)
	} else if len(s2) == 1 {
		result = append(result, string(s1[0])+s2)
		combi(s1[1:], s2, result)
	} else {
		result = append(result, string(s1[0])+string(s2[0]))
		combi(s1[1:], s2[1:], result)
	}
}

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	result:=""
	i:=0
	for ;; {
		c:=""
		if i< len(strs[0]){
			c=string(strs[0][i])
		} else {
			return result
		}
		for _, str:=range strs[1:] {
			if i>=len(str) || c!=string(str[i]) {
				return result
			}
		}
		result += c
		i++
	}
	return result
}

func main()  {
	fmt.Println(longestCommonPrefix([]string{
		"abc","bc","a",
	}))
	fmt.Println(longestCommonPrefix([]string{
		"bc","bcdasdfasd","bcdd",
	}))
}

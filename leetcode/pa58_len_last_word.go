package main

import "fmt"

func lengthOfLastWord(s string) int {
	l:=len(s)
	if l == 0 {
		return 0
	}
	len:=0
	begin:=false
	for i:=l-1;i>=0;i--{
		r:=s[i]
		if r == ' ' && begin {
			return len
		}
		if r!= ' ' {
			begin = true
			len++
		}
	}
	return len
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("   "))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("adsfasdfasd"))
	fmt.Println(lengthOfLastWord("adsfasdfasd "))
}
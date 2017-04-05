package main

import "fmt"

func isValid(s string) bool {
	var pop rune
	var t []rune
	for _, r := range s {
		if len(t) == 0 {
			t = append(t, r)
			continue
		}
		pop = t[len(t)-1]
		if (r == '}' && pop =='{') || (r == ')' && pop =='(') || (r == ']' && pop =='['){
			t = t[:len(t)-1]
		} else {
			t = append(t, r)
		}
	}
	return len(t) == 0
}

func main() {
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[)]}"))
	fmt.Println(isValid("{[()]}"))
	fmt.Println(isValid("{[()}]"))
}

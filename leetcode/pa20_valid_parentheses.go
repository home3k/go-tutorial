package main

import "fmt"

func isValid(s string) bool {
	var p1, p2, p3 int
	for _, c := range s {
		if c == '(' {
			p1 += 1
		}
		if c == ')' {
			p1 -= 1
		}
		if c == '{' {
			p2 += 1
		}
		if c == '}' {
			p2 -= 1
		}
		if c == '[' {
			p3 += 1
		}
		if c == ']' {
			p3 -= 1
		}
	}
	return p1 == 0 && p2 ==0 && p3==0
}

func main()  {
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[)]}"))
	fmt.Println(isValid("{[()]}"))
	fmt.Println(isValid("{[()}]"))
}

package main

import "fmt"

//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	var result []string
	var contain = make(map[string]bool)
	var s string
	for _, p:=range generateParenthesis(n-1) {
		s = "("+ p +")"
		if _, ok :=contain[s];!ok {
			result = append(result, s)
			contain[s] = true
		}
		s = p + "()"
		if _, ok :=contain[s];!ok {
			result = append(result, s)
			contain[s] = true
		}
		s = "()" + p
		if _, ok :=contain[s];!ok {
			result = append(result, s)
			contain[s] = true
		}
	}
	return result
}

func main()  {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(0))
}

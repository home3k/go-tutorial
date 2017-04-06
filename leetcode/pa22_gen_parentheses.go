package main

import "fmt"

//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
// (  )
// x x x | x x x
// xx ( xxx
// x((xxx
// (((xxx
//
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var result []string = []string{}
	gen(&result, "", n, n)
	return result
}

func gen(result *[]string, p string, left int, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, p)
	}
	if left > 0 {
		gen(result, p+"(", left-1, right)
	}
	if right > 0 && right>left{
		gen(result, p+")", left, right-1)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(0))
}

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	ne := 0
	t :=x
	for x > 0 {
		ne = ne*10 + x%10
		x = x / 10
	}
	return ne == t

}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-1123))
	fmt.Println(isPalindrome(0))
}

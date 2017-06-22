package main

import "fmt"

/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.

 */
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	sub:=1
	for i:=len(digits)-1; i>=0; i-- {
		v:=digits[i]
		next:=v+sub
		if next>=10 {
			sub=1
			next -= 10
		} else {
			sub=0
		}
		digits[i]=next
	}
	if sub >0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main()  {
	fmt.Println(plusOne([]int{1,2,0,2}))
	fmt.Println(plusOne([]int{1,2,9,9}))
	fmt.Println(plusOne([]int{9,9,9,9}))
}

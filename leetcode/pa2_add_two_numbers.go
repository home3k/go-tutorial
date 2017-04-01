package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	for l1 == nil && l2 == nil {
		return nil
	}

	result := new(ListNode)
	current := result
	// init Val
	val := 0
	for l1 != nil && l2 != nil {
		current.Next = new(ListNode)
		current = current.Next

		value := l1.Val + l2.Val + val
		val = value / 10
		current.Val = value % 10

		l1 = l1.Next
		l2 = l2.Next

	}

	for l1 != nil {
		current, l1, val = processNode(current, l1, val)
	}
	for l2 != nil {
		current, l2, val = processNode(current, l2, val)
	}

	if val > 0 {
		current.Next = new(ListNode)
		current = current.Next
		current.Val = val
	}
	return result.Next
}

func processNode(current *ListNode, l *ListNode, val int) (*ListNode, *ListNode, int) {
	current.Next = new(ListNode)
	current = current.Next
	value := l.Val + val
	val = value / 10
	current.Val = value % 10
	l = l.Next
	return current, l, val
}

func print(l *ListNode)  {
	fmt.Print("[")
	t := l
	for t!=nil {
		fmt.Printf(" %d ",t.Val)
		t = t.Next
	}
	fmt.Print("]\n")
}

func build(vals []int) *ListNode {
	result := new(ListNode)
	current := result
	for index, val := range vals {
		current.Val = val
		if index != len(vals) -1 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}
	return result
}
func main()  {
	l1 := build([]int{2,4,3,9})
	l2 := build([]int{5,6,4,9,9})
	result := addTwoNumbers(l1, l2)
	print(l1)
	print(l2)
	print(result)
}


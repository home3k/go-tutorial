package main

import (
	"fmt"
)

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first, second:= head, head.Next
	pointer:=new(ListNode)
	pointer.Next = first
	result:=pointer
	for first!=nil && second!=nil  {
		pointer.Next = second
		first.Next = second.Next
		second.Next = first
		pointer = first
		first = pointer.Next
		if first!=nil {
			second = first.Next
		}
	}
	return (*result).Next
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

/*
2 -> 1 ->3 ->4

head

pointer.next=second
first.next=second.next
second.next=first
first
second

 */

func main()  {
	l := build([]int{1,2,3,4})
	print(l)
	swapPairs(l)
	print(l)
}

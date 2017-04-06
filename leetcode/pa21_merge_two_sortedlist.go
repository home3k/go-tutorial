package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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


// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2==nil {
		return l1
	}
	var head *ListNode
	if l1.Val>l2.Val {
		// 3, 4 ,5   1 2 3
		head = l2
		head.Next = mergeTwoLists(l1, l2.Next)
	} else {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	}
	return head
}

func main()  {
	l1 := build([]int{1,2,3,9})
	l2 := build([]int{5,6,4,9,9})
	r:=mergeTwoLists(l1, l2)
	print(l1)
	print(l2)
	print(r)
}
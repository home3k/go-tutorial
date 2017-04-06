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
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1!=nil {
		if l2 == nil {
			return l1
		}
	} else {
		return l2
	}
	head:=l1
	if l1.Val>l2.Val {
		head=l2
	}
	var p *ListNode
	v1, v2:=l1.Val, l2.Val
	for l1!=nil && l2!=nil {
		if l1.Val>v2 {
			if l1 < v2 {
				
			}	
		} else {
		}
	}
	if l1!=nil {
		l2.Next = l1
	}
	if l2!=nil {
		l1.Next = l2
	}
	if l1f {
		return h1
	} else {
		return h2
	}
}



func main()  {
	l1 := build([]int{1,2,3,9})
	l2 := build([]int{5,6,4,9,9})
	r:=mergeTwoLists(l1, l2)
	print(l1)
	print(l2)
	print(r)
}

1 2 3 9

5 6 4 9 9

1
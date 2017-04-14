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
/*
   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.

*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	result := head
	t := []*ListNode{}
	for head != nil {
		t = append(t, head)
		count++
		head = head.Next
	}
	if n >= count {
		if count >1 {
			return t[1]
		} else {
			return nil
		}
	}
	if n <= 1 {
		t[count-n-1].Next = nil
	} else {
		t[count-n-1].Next = t[count-n+1]
	}
	return result
}

func print(l *ListNode) {
	fmt.Print("[")
	t := l
	for t != nil {
		fmt.Printf(" %d ", t.Val)
		t = t.Next
	}
	fmt.Print("]\n")
}

func build(vals []int) *ListNode {
	result := new(ListNode)
	current := result
	for index, val := range vals {
		current.Val = val
		if index != len(vals)-1 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}
	return result
}
func main() {
	print(removeNthFromEnd(build([]int{2, 4, 3, 9}), 2))
	print(removeNthFromEnd(build([]int{5, 6, 4, 9, 9}), 3))
	print(removeNthFromEnd(build([]int{1}), 1))
	print(removeNthFromEnd(build([]int{1, 2}), 1))
	print(removeNthFromEnd(build([]int{1, 2}), 2))
}

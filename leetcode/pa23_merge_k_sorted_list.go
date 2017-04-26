package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

1 2 3 4
2 3 5 6

1 2 3 4  2 3 5 6 3 6 9
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	if l == 2 {
		l1 := lists[0]
		l2 := lists[1]
		if l1 != nil && l2 != nil {
			result := l2
			b1, b2 := l2, l1
			if l1.Val < l2.Val {
				result = l1
				b1, b2 = l1, l2
			}
			for b1 != nil && b2 != nil {
				if b1.Next != nil && b1.Next.Val > b2.Val {
					t := b1.Next
					b1.Next = b2
					b2 = t
				} else if b1.Next == nil {
					b1.Next = b2
					b2 = nil
				} else {
					b1 = b1.Next
				}
			}
			return result
		} else if l1 != nil {
			return l1
		} else {
			return l2
		}
	} else {
		return mergeKLists([]*ListNode{lists[0], mergeKLists(lists[1:])})
	}
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
	l1 := build([]int{1})
	l2 := build([]int{0})
	print(mergeKLists([]*ListNode{l1, l2}))
}

package main

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k < 2 {
		return head
	}
	start, end, reverse := rev(head, k)
	result := start
	for reverse && end != nil && end.Next != nil {
		ns, ne, nr := rev(end.Next, k)
		end.Next = ns
		end = ne
		reverse = nr
	}
	return result
}

func rev(head *ListNode, k int) (start, end *ListNode, reverse bool) {
	if k == 1 {
		if head != nil {
			return head, head, true
		} else {
			return head, nil, false
		}
	} else {
		// k > 1
		if head == nil || head.Next == nil {
			return head, nil, false
		}
		// head !=nil ^^ head.next!=nil
		n := head.Next
		nextNode, lastNode, reverse := rev(n, k-1)
		if reverse {
			head.Next = lastNode.Next
			lastNode.Next = head
			return nextNode, head, true
		} else {
			return head, nil, false
		}
	}
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
	l1 := build([]int{1, 2, 3, 4, 5, 6})
	reverseKGroup(l1, 3)
}

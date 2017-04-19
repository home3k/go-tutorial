package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil && q==nil  {
		return true
	}
	if p==nil || q==nil {
		return false
	}
	l1, l2, r1, r2:=p.Left,q.Left, p.Right, q.Right
	if l1!=nil && l2!=nil {
		if !isSameTree(l1, l2) {
			return false
		}
	} else if !(l1==nil && l2==nil) {
		return false
	}
	if r1!=nil && r2!=nil {
		if !isSameTree(r1, r2) {
			return false
		}
	} else if !(r1==nil && r2==nil) {
		return false
	}
	return p.Val == q.Val
}

func main() {

}

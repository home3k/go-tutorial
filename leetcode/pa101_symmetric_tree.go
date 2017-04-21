package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

    1
   / \
  2   2
 / \ / \
3  4 4  3

3 2 4 1 4 2 3

 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right:=root.Left, root.Right
	return _isSymmetric(left, right)
}

func _isSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else {
		if left.Val != right.Val {
			return false
		} else {
			return _isSymmetric(left.Right, right.Left) && _isSymmetric(left.Left, right.Right)
		}
	}
}

package main

import "math"

var preVal int = math.MinInt

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !IsValidBST(root.Left) {
		return false
	}
	if root.Val < preVal {
		return false
	}
	preVal = root.Val
	return IsValidBST(root.Right)
}

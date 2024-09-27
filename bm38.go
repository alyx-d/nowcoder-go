package main

func LowestCommonAncestor2(root *TreeNode, p int, q int) int {
	if root == nil {
		return -1
	}
	if root.Val == p || root.Val == q {
		return root.Val
	}
	left := LowestCommonAncestor2(root.Left, p, q)
	right := LowestCommonAncestor2(root.Right, p, q)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	return root.Val
}

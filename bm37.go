package main

func LowestCommonAncestor(root *TreeNode, p int, q int) int {
	if root == nil {
		return -1
	}
	if p < root.Val && q < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}
	if p > root.Val && q > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}
	return root.Val
}

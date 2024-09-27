package main

func Mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := Mirror(root.Left)
	right := Mirror(root.Right)
	root.Left = right
	root.Right = left
	return root
}

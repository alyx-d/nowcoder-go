package main

var pre *TreeNode
var head *TreeNode

func Convert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	Convert(root.Left)
	if pre == nil {
		pre = root
		head = root
	} else {
		pre.Right = root
		root.Left = pre
		pre = root
	}
	Convert(root.Right)
	return head
}

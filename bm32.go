package main

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := new(TreeNode)
	root.Val = t1.Val + t2.Val
	root.Left = MergeTrees(t1.Left, t2.Left)
	root.Right = MergeTrees(t1.Right, t2.Right)
	return root
}

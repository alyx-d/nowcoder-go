package main

func InorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, InorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, InorderTraversal(root.Right)...)
	return ans
}

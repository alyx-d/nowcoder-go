package main

func PostorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, PostorderTraversal(root.Left)...)
	ans = append(ans, PostorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

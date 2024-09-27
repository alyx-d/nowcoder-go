package main

func PreorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, PreorderTraversal(root.Left)...)
	ans = append(ans, PreorderTraversal(root.Right)...)
	return ans
}

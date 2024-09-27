package main

func IsBalanced_Solution(pRoot *TreeNode) bool {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if abs(left-right) > 1 || left == -1 || right == -1 {
			return -1
		}
		return max(left, right) + 1
	}
	return dfs(pRoot) != -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

package main

func ReConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	var dfs func([]int, int, int, []int, int, int) *TreeNode
	dfs = func(p []int, ps int, pe int, v []int, vs int, ve int) *TreeNode {
		if ps > pe || vs > ve {
			return nil
		}
		root := new(TreeNode)
		root.Val = p[ps]
		for i := vs; i < len(vinOrder); i++ {
			if vinOrder[i] == p[ps] {
				root.Left =
					dfs(p, ps+1, ps+i-vs, v, vs, i-1)
				root.Right =
					dfs(p, ps+i-vs+1, pe, v, i+1, ve)
				break
			}
		}
		return root
	}
	return dfs(preOrder, 0, len(preOrder)-1, vinOrder, 0, len(vinOrder)-1)
}

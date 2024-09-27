package main

import "container/list"

func Solve(preOrder []int, vinOrder []int) []int {
	var dfs func([]int, int, int, []int, int, int) *TreeNode
	dfs = func(p []int, ps, pe int, v []int, vs, ve int) *TreeNode {
		if ps > pe || vs > ve {
			return nil
		}
		root := new(TreeNode)
		root.Val = p[ps]
		for i := vs; i < len(v); i++ {
			if v[i] == p[ps] {
				root.Left =
					dfs(p, ps+1, ps+i-vs, v, vs, i-1)
				root.Right =
					dfs(p, ps+i-vs+1, pe, v, i+1, ve)
				break
			}
		}
		return root
	}
	root := dfs(preOrder, 0, len(preOrder)-1, vinOrder, 0, len(vinOrder)-1)
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for size > 0 {
			ele := queue.Front()
			queue.Remove(ele)
			node := ele.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			size--
			if size == 0 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}

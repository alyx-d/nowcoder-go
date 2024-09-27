package main

import (
	"strconv"
	"strings"
)

func Serialize(root *TreeNode) string {
	s := ""
	var dfs func(*TreeNode)
	dfs = func(t *TreeNode) {
		if t == nil {
			s += "#,"
			return
		}
		s += strconv.Itoa(t.Val) + ","
		dfs(t.Left)
		dfs(t.Right)
	}
	dfs(root)
	return s
}

func Deserialize(s string) *TreeNode {
	str := s
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(str) == 0 {
			return nil
		}
		var node *TreeNode
		if str[0] != '#' {
			node = new(TreeNode)
			idx := strings.Index(str, ",")
			node.Val, _ = strconv.Atoi(str[:idx])
			str = str[idx+1:]
			node.Left = dfs()
			node.Right = dfs()
		} else {
			str = str[2:]
		}
		return node
	}
	return dfs()
}

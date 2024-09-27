package main

import "container/list"

func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root)
	flag := false
	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*TreeNode)
		queue.Remove(element)
		if node == nil {
			flag = true
			continue
		}
		if flag {
			return false
		}
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
	return true
}

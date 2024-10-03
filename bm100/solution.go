package bm100

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type Solution struct {
	capacity, size int
	dict           map[int]*Node
	head, tail     *Node
}

func Constructor(c int) Solution {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-1, -1, nil, nil}
	head.Next = tail
	tail.Prev = head
	return Solution{c, 0, make(map[int]*Node), head, tail}
}

func (s *Solution) insertHead(node *Node) {
	node.Prev = s.head
	node.Next = s.head.Next
	s.head.Next.Prev = node
	s.head.Next = node
}

func (s *Solution) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (s *Solution) moveToHead(node *Node) {
	s.removeNode(node)
	s.insertHead(node)
}

func (s *Solution) removeTail() *Node {
	node := s.tail.Prev
	s.removeNode(node)
	return node
}

func (s *Solution) get(key int) int {
	node, ok := s.dict[key]
	if !ok {
		return -1
	}
	s.moveToHead(node)
	return node.Val
}

func (s *Solution) set(key, val int) {
	node, ok := s.dict[key]
	if !ok {
		node = &Node{key, val, nil, nil}
		s.dict[key] = node
		s.insertHead(node)
		s.size++
		if s.size > s.capacity {
			delete(s.dict, s.removeTail().Key)
			s.size--
		}
	} else {
		node.Val = val
		s.moveToHead(node)
	}
}

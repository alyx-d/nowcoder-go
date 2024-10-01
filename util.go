package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	Start, End int
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func ReadCase(f func(*ListNode)) {
	if len(os.Args) < 2 {
		return
	}
	test_case := os.Args[1]
	fd, err := os.Open(test_case)
	if err != nil {
		fmt.Println(err)
		return
	}
	fs := bufio.NewScanner(fd)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		// every line
		line := fs.Text()
		root := ParseToListNode(line)
		f(root)
	}
	defer fd.Close()
}

func ParseToListNode(line string) *ListNode {
	if len(line) <= 2 {
		return nil
	}
	strs := strings.Split(line[1:len(line)-1], ",")
	root := new(ListNode)
	curr := root
	for _, val := range strs {
		num, _ := strconv.Atoi(val)
		curr.Next = new(ListNode)
		curr.Next.Val = num
		curr = curr.Next
	}
	return root.Next
}

func (root *ListNode) ToString() string {
	curr := root
	res := "{"
	seq := ""
	for curr != nil {
		res += seq + strconv.Itoa(curr.Val)
		seq = ","
		curr = curr.Next
	}
	res += "}"
	return res
}

type Plus interface {
	Add() Plus
}

func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

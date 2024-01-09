package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	right := head
	tmp := head
	left := reverse(tmp)

	for left != nil || right != nil {
		if left == right {
			left = left.Next
			right = right.Next
		} else {
			return false
		}
	}
	return true
}

func reverse(h *ListNode) *ListNode {
	var par *ListNode
	cur := h
	for cur != nil {
		temp := cur.Next
		cur.Next = par
		par = cur
		cur = temp
	}
	return par
}
func reverse2(h *ListNode) *ListNode {
	par := &ListNode{Next: nil}
	cur := h
	for cur != nil {
		tmp := cur.Next
		cur.Next = par.Next
		par.Next = cur
		cur = tmp
	}
	return par.Next
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d
	x := a

	tmp := reverse2(x)
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	for a != nil {
		fmt.Print(a.Val)
		a = a.Next
	}
	for d != nil {
		fmt.Print(d.Val)
		d = d.Next
	}
}

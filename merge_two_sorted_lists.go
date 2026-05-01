package main

import "fmt"
import "slices"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr []int

	for l1 != nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		arr = append(arr, l2.Val)
		l2 = l2.Next
	}

	slices.Sort(arr)

	dummy := &ListNode{}
	current := dummy

	for _, v := range arr {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return dummy.Next
}

func main() {
	listA := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	
	listB := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	result := mergeTwoLists(listA, listB)

	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
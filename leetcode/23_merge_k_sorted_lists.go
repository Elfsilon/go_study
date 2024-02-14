package leetcode

import (
	"fmt"
	"strings"
)

// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	lst := make([]string, 0)
	for {
		lst = append(lst, fmt.Sprint(n.Val))

		if n.Next == nil {
			break
		} else {
			n = n.Next
		}
	}
	return "[" + strings.Join(lst, ",") + "]"
}

func Append(n *ListNode, el int) *ListNode {
	if n == nil {
		return &ListNode{
			Val: el,
		}
	}

	curr := n
	var prev *ListNode

	for {
		if curr == nil {
			prev.Next = &ListNode{
				Val: el,
			}
			break
		}

		if el < curr.Val {
			if prev == nil {
				n = &ListNode{
					Val:  el,
					Next: curr,
				}
				break
			}

			prev.Next = &ListNode{
				Val:  el,
				Next: curr,
			}
			break
		}

		prev = curr
		curr = curr.Next
	}

	return n
}

func MergeKLists(lists []*ListNode) *ListNode {
	fmt.Println(lists)
	if len(lists) == 0 {
		return nil
	}

	res := lists[0]

	for _, head := range lists[1:] {
		curr := head
		for curr != nil {
			res = Append(res, curr.Val)
			curr = curr.Next
		}
	}

	return res
}

package _3_merge_k_link_lisk

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]
 

提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	newList := make([]*ListNode, 0)
	for _, v := range lists {
		if v != nil {
			newList = append(newList, v)
		}
	}
	lists = newList
	if len(lists) == 0 {
		return nil
	}
	var head, tail *ListNode
	for len(lists) > 0 {
		lists = minHeap(lists)
		if lists[0] == nil {
			lists = lists[1:]
			continue
		}
		if head == nil {
			head = &ListNode{Val: lists[0].Val}
			tail = head
		} else {
			tail.Next = &ListNode{Val: lists[0].Val}
			tail = tail.Next
		}
		if lists[0].Next != nil {
			lists[0] = lists[0].Next
		} else {
			lists = lists[1:]
		}
		if len(lists) == 0 {
			break
		}
	}
	return head
}

func minHeap(lists []*ListNode) []*ListNode {
	N := len(lists) - 1
	for k := N / 2; k >= 0; k-- {
		sink(lists, k, N)
	}
	return lists
}

func sink(s []*ListNode, k, N int) {
	for {
		i := 2*k + 1
		if i > N {
			break
		}
		if i < N && s[i+1].Val < s[i].Val {
			i++
		}
		if s[k].Val <= s[i].Val {
			break
		}
		s[i], s[k] = s[k], s[i]
		k = i
	}
}

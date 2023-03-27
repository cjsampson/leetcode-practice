package merge_two_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	for l1.Val != nil && l2.Val != nil {
//		if l1.Val < l2.Val {
//			node.Next = mergeTwoLists(l1, l2)
//			return l1
//		}
//	}
//	return nil
//}

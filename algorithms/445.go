package algorithms

func AddTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseListNode(l1)
	l2 = reverseListNode(l2)
	return reverseListNode(AddTwoNumbers2(l1, l2))
}

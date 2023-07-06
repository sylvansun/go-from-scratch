package algorithms

func AddTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	return reverseList(AddTwoNumbers2(l1, l2))
}

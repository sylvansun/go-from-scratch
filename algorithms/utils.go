package algorithms

func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func squareDistanceFloat64(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

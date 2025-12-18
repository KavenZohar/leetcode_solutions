/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := ListNode{}
	m := &n
	a := 0
	for l1 != nil || l2 != nil || a != 0 {
		s := a
		if l1 != nil {
			s = s + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s = s + l2.Val
			l2 = l2.Next
		}
		a = s / 10
		m.Next = &ListNode{ Val: s % 10 }
        m = m.Next
	}
	return n.Next
}

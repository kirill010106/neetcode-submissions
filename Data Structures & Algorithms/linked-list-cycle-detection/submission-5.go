/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    slow, fast := head, head
	if head == nil {
		return false
	}
	for {
		if slow.Next != nil {
			slow = slow.Next
		} else {
			return false
		}
		if fast.Next != nil && fast.Next.Next != nil{
			fast = fast.Next.Next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}
}

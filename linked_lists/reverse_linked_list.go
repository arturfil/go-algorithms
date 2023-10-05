package linkedlists

func reverseList(head *ListNode) *ListNode {
    var prev, temp *ListNode

    for head != nil {
        temp = head.Next
        head.Next = prev
        prev = head
        head = temp
    }

    return prev
}

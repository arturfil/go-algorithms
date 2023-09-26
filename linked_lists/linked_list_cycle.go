package linkedlists

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil { 
        return false 
    }
    p1, p2 := head, head.Next  

    for p1 != p2 {
        if p2 == nil || p2.Next == nil { 
            return false
        }

        p1, p2 = p1.Next, p2.Next.Next
    }

    return p1 == p2
}

/*
    Testing
    -
    head := &linkedlists.ListNode{Val: 3, Next: nil}
    n2 := &linkedlists.ListNode{Val: 2, Next: nil}
    n3 := &linkedlists.ListNode{Val: 0, Next: nil}
    n4 := &linkedlists.ListNode{Val: -4, Next: nil}
    
    head.Next = n2
    n2.Next = n3
    n3.Next = n4

    fmt.Println(head)

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    splitter := dummy
    curr := dummy

    for curr.Next != nil {
        if curr.Next.Val < x {
            if splitter == curr {
                splitter = splitter.Next
                curr = curr.Next
                continue
            }
            temp := curr.Next
            curr.Next = curr.Next.Next

            temp.Next = splitter.Next
            splitter.Next = temp
            splitter = splitter.Next
            continue
        }
        curr = curr.Next
    }

    return dummy.Next
}
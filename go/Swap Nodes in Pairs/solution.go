/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    
    if head == nil {
        return nil
    }
    
    var newHead *ListNode = nil
    if head.Next == nil {
        newHead = head
    } else {
        
        var cur *ListNode = head
        var last *ListNode = nil

        for cur != nil {
            second := cur.Next
            if second == nil {
                break
            }
            rest := second.Next
            second.Next = cur
            cur.Next = rest
            
            if last != nil {
                last.Next = second
            } else {
                newHead = second
            }
            last = cur
            
            // Advance
            cur = cur.Next

        }       
    }
    
   

    return newHead
}
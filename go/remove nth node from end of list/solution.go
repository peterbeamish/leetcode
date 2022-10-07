/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    
    arr := make([]*ListNode,0)
    arr = append(arr, head)
    
    cur := head
    
    for cur.Next != nil {
        arr = append(arr, cur.Next)
        cur = cur.Next
    }
    
    return deleteNth(arr, n)
}


func deleteNth(arr []*ListNode, n int) *ListNode {
    if arr == nil || len(arr) == 0 {
        return nil
    }
    
    if n> len(arr){
        // Handle error and return
        return arr[0]
    }
    
    var last,  next *ListNode
    
    indexOfLast := (len(arr)-n)-1
    //indexOfTarget := (len(arr)-n)
    indexOfNext := (len(arr)-n)+1
    
    if indexOfLast<0 {
        // Handle remove head
        // MUST RETURN
        return arr[0].Next
    } else {
        last = arr[indexOfLast]
    }
    
    //target = arr[indexOfTarget]
    
    if indexOfNext <= (len(arr)-1){
        next = arr[indexOfNext]
    }
    
    // last must be not nil by here
    if last == nil {
        fmt.Println("Last is nil. This shouldn't happen")
        return nil
        // TODO HANDLE
    }
    
    // Cut out the target
    last.Next = next
    
    return arr[0]
    
}
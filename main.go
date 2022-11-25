
package main
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 
 * @param head ListNode类 the head
 * @return bool布尔型
*/
func re (head *ListNode) *ListNode{
    p := new(ListNode)
    for head != nil{
        t := head.Next
        head.Next = p
        head = t
    }
    return p
    
}
func isPail( head *ListNode ) bool {
    if head == nil{
        return true
    }
    // write code here
    fast := head
    low := head
    for fast != nil && fast.Next != nil{
            fast = fast.Next.Next
            low = low.Next
    }
    // fast at the end
    // low at mid or mid + 1
    low = re(low)
    fast = head
    for low != nil {
        if low.Val != fast.Val {
            return false
        }
        fast = fast.Next
        low = low.Next

    }
    return true
}

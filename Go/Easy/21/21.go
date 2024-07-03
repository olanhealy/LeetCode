 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    solution := &ListNode{}
    temp := solution

    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {
            temp.Next = list2
            list2 = list2.Next
        } else {
            temp.Next = list1
            list1 = list1.Next
        }
        temp = temp.Next
    }

    if (list1 == nil) {
        temp.Next = list2
    }
    if list2 == nil {
        temp.Next = list1
    }
    return solution.Next

}

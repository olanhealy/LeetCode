/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

    if root == nil {
        return []int {}
    } 

    result := []int{}
    stack := []*TreeNode{}
    current := root

    for current != nil || len(stack) > 0 {

    for current != nil {
        stack = append(stack, current)
        current = current.Left
    }

    current = stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    result = append(result, current.Val)

    current = current.Right

    }
    return result

    

    
}

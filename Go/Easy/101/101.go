/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

    if root == nil {
        return true
    }

    var stack = []*TreeNode{}
    stack = append(stack, root.Left)
    stack = append(stack, root.Right)

    for len(stack) > 0  {
        right := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        left := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if left == nil && right == nil {
            continue
        }

        if left == nil || right == nil {
            return false
        }

        if left.Val != right.Val {
            return false
        }

        stack = append(stack, left.Left)
        stack = append(stack, right.Right)
        stack = append(stack, left.Right)
        stack = append(stack, right.Left)
    }
    return true

    
}

# Notes
Did few days ago forgot to push
# Complexity
- O(N)



# Code
```Go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {

    if root == nil {
        return []int{}
    }

    result := []int{}
    stack1 := []*TreeNode{}
    stack2 := []*TreeNode{}
    current := root
    stack1 = append(stack1, current)

    for len(stack1) > 0 {
        current = stack1[len(stack1) - 1]
        stack1 = stack1[:len(stack1) -1 ]
        stack2 = append(stack2, current)
        

        if (current.Left != nil) {
            stack1 = append(stack1, current.Left)
        }

        if (current.Right != nil) {
            stack1 = append(stack1, current.Right)
        }
    }

        for len(stack2) > 0 {
            current := stack2[len(stack2) - 1]
            stack2 = stack2[:len(stack2) -1 ]
            result = append(result, current.Val)
            
        }

    
    return result
    
}
```

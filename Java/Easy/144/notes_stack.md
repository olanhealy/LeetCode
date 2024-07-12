# Notes
this is using stack. Had to look up  https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

once I kinda got my head around it it's actually very straight forward
# Complexity
O(n)

# Code
```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        if (root == null) {
        return new ArrayList<Integer>();
        
    }

    ArrayList<Integer> result = new ArrayList<>();
    Stack<TreeNode> stack = new Stack<>();

    stack.push(root);
    TreeNode node;
    
    while (!stack.isEmpty()) {
        node = stack.pop();
        result.add(node.val);
        if (node.right!=null) {
            stack.push(node.right);
        }
        if (node.left!=null) {
            stack.push(node.left);
        }
    }
    return result;
    }

}

```

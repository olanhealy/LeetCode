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
    public List<Integer> postorderTraversal(TreeNode root) {

    if (root == null ) {
        return new ArrayList<Integer>();
    }

    ArrayList<Integer> result = new ArrayList<>();
    Stack<TreeNode> stack1 = new Stack<>();
    Stack<TreeNode> stack2 = new Stack<>();

    stack1.push(root);
    TreeNode node;

    while (!stack1.isEmpty() ) {
        node = stack1.pop();
        stack2.push(node);

        if (node.left != null) {
            stack1.push(node.left);
        }

        if (node.right != null) {
            stack1.push(node.right);
        }
    }
        while (!stack2.isEmpty()) {
            node = stack2.pop();
            result.add(node.val);
        }
        
    
    return result;
    }
}

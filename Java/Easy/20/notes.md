# Notes
Was first trying to convert string to char array and use a .contains and check if next char is the matching parenthesis but wouldnt work
went to discussion and seen people talk about stack so I had tos search up how to use stack


# Complexity

O(N) as for loop iterates over each character in string


# Code
```Java
import java.util.Stack;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<Character>();
        //s = "(){}"

        for (char c : s.toCharArray()) {
            if (c == '(' || c == '[' || c == '{') {
                stack.push(c);
            } else {
                if ( stack.isEmpty()) {
                    return false;
                }
                char inStack = stack.peek();
                if ((c == ')' && inStack == '(') || (c == ']' && inStack == '[') || (c == '}' && inStack == '{')) {
                    stack.pop();
                } else {
                    return false;
                }
            }
        }
        return stack.isEmpty();
    }
}
```

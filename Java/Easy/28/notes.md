# Notes
Solved within 5 minutes very easy and straight forward

# Complexity
O(m x n) as indexOf is m x n runtime

# Code
```Java
class Solution {
    public int strStr(String haystack, String needle) {

     

        
        if (haystack.contains(needle)) {
            return haystack.indexOf(needle);
        }
        return -1;
    }

}
```

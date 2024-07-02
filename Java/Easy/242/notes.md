# Notes
Had the idea from the start which was comparing the strings alphabetically. https://www.quora.com/How-do-you-check-if-two-char-arrays-are-equal-in-Java-without-a-method#:~:text=You%20can%20use%20Arrays.,that%20you%20can%20return%20true.

# Complexity
> O(n log (n)) which is in the Arrays.sort

# Code
```Java
class Solution {
    public static boolean isAnagram(String s, String t) {

        char sNew[] = s.toCharArray();
        char tNew[] = t.toCharArray();

        if (sNew.length != tNew.length) {
            return false;
        }
        Arrays.sort(tNew);
        Arrays.sort(sNew);


        return Arrays.equals(sNew, tNew);
    }
}

```

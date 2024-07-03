# Notes
had the idea from the start, one issue I did was for some reason I was adding int x and the j hoping it would equal the target but it just broke the loop


# Complexity
> O(n^2) as it is a nested for loop


# Code
```java
class Solution {
    public int[] twoSum(int[] nums, int target) {

        for (int i = 0; i < nums.length; i++) {
            int x = target - nums[i];
            
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] == x) {
                    return new int[] { i, j };
                }
            }
        }

        return null;
    }
}


```

# Notes
So did it with a hashmap after as it had a better run time. Need to get used to trying to do the optimal way first

# Complexity
> O(n) as it only loops over array once

# Code
```
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> solution = new HashMap<>();
        
        for (int i = 0; i < nums.length; i++) {
            int x = target - nums[i];
            if (solution.containsKey(x)) {
                return new int[] { solution.get(x), i};
            }
            solution.put(nums[i], i);
        }
        
        return null;

    }
}


```

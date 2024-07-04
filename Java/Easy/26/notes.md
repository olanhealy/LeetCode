# Notes
very easy todo this way. just check if nothing in array is one element so will have 0 non duplicates
then we can start with seeing if element at int i = 1 equals one previous
if they dont, then can add this element to array and increment non dupe

# Complexity
> O(n log n) has Arrays.sort has tat runtime, 

# Code
```Java
class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) return 0; // so we can use for loop properly

        Arrays.sort(nums);

        int nonDuplicate = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[i -1]) {
                nums[nonDuplicate] = nums[i];
                nonDuplicate++;
                
            }
        }
    return nonDuplicate;
    }
}
```

# Notes
Terribly worded question. It says it wants you to return new array but then only accepts and int as return it makes no sense but anyways had to look at discussion first

Decided to use hashmap but had to search the for loop part as I was doing 
~~~
for (int i: nums) 
~~~
but it was missing elements 

# Complexity

> O(n) as only has to go through for loop



# Code
```
class Solution {
    public int removeDuplicates(int[] nums) {

        Map<Integer, Integer> map = new HashMap<>();

        int notDuplicate = 0;

        for (int i = 0; i < nums.length; i++) {
            if (!map.containsKey(nums[i])) {
                map.put(nums[i], 1);
                nums[notDuplicate] = nums[i];
                notDuplicate++;
            }

        }
        return notDuplicate;
    }
}

```

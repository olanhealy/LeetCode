# Notes

this is a hashmap answer. had to use https://www.w3schools.com/java/java_hashmap.asp
to give refresher on hashmaps in java

# Complexity
O(N) as it only looks at each number once, adds to hashmap then if number comes up again it will end the for loop

# Code
```
class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Boolean> place = new HashMap<>();

        for (int i : nums) {
            if (place.containsKey(i)) {

                return true;
            }

            place.put(i, true);
        }
        return false;
    }
}
    
```

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
    


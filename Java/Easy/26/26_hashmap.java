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


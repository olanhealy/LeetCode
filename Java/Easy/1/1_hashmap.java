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



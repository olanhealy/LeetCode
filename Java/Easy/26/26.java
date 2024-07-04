 {
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


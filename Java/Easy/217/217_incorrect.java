class Solution {
    public boolean containsDuplicate(int[] nums) {

        int i ;
        int j ;

        for (i = 0; i < nums.length ; i++) {
            for (j = 0; j < nums.length ; j++) {
                if (nums[i] == nums[j] ) {
                    return true;
                }

            }


        }
        return false;
    }
}


class Solution {
    public boolean containsDuplicate(int[] nums) {

        int i, j ;

        for (i = 0; i < nums.length ; i++) {
            int place = nums [i];
            for (j = i + 1; j < nums.length ; j++) {
                if (nums[j] == place ) {
                    return true;
                }

            }
        }
        return false;
    }
}
